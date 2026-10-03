// register 将新 CRUD 模块的登记代码一次性注入手写装配文件,替代原 Wire 时代的
// "改 wire set + make wire" 自动化。注入点为各文件中的 register:* 锚点注释:
//
//	领域服务(BFF 无仓储形态亦支持): cmd/server/wiring.go            仓储(服务客户端)行 / 服务行 / 服务实参
//	                                   internal/server/{grpc,rest}*.go 服务形参 / 路由注册调用
//
// 仅覆盖标准形态(data.New<Ent>Repo(ctx, entClient) /
// service.New<Ent>Service(ctx, <ent>Repo);BFF 形态为 data/client 的
// New<Ent>ServiceClient 与对应 HTTP 注册);依赖更多的模块请手工调整注入行。
//
// 用法:
//
//	go run ./tools/register -entity sprint -svc kanban          # 仓储+服务+gRPC 注册
//	go run ./tools/register -entity me -svc front               # BFF 客户端+服务+REST 注册
//	go run ./tools/register -entity collection -svc app -pkg contentV1   # 指定路由包别名
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

type patch struct {
	path    string   // 目标文件
	anchors []string // [锚点行, 注入行] 成对;命中首个锚点后,在其后插入对应注入行
	skipIf  string   // 文件已包含该片段(多值用 | 分隔)时跳过(幂等)
}

var reRouteAlias = regexp.MustCompile(`([A-Za-z0-9_]+)\.Register[A-Za-z0-9_]*Service(?:Server|HTTPServer)\(`)

func main() {
	entity := flag.String("entity", "", "模块名:sprint / dict_entry / dictEntry 均可")
	svc := flag.String("svc", "", "目标服务目录名:kanban / front / dict ...")
	pkg := flag.String("pkg", "", "路由包别名;缺省从既有 Register 调用自动探测")
	flag.Parse()

	if *entity == "" || *svc == "" {
		flag.Usage()
		os.Exit(2)
	}
	base := filepath.Join("app", *svc, "service")
	if _, err := os.Stat(base); err != nil {
		fmt.Fprintf(os.Stderr, "register: 服务目录 %s 不存在(-svc 取 app/<名字> 的 <名字>)\n", base)
		os.Exit(2)
	}

	typ := pascal(*entity)
	name := lowerName(typ)
	wiringPath := filepath.Join(base, "cmd", "server", "wiring.go")

	// 服务端文件:含 register:{grpc,rest}-param 锚点的那个
	serverPath, flavor, err := findServerFile(base)
	if err != nil {
		fmt.Fprintln(os.Stderr, "register:", err)
		os.Exit(1)
	}

	// 路由包别名:优先 -pkg,否则从既有 Register 调用探测
	routePkg := *pkg
	if routePkg == "" {
		routePkg, err = detectRoutePkg(serverPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "register:", err)
			os.Exit(1)
		}
	}

	registerFn := "Register" + typ + "ServiceServer"
	if flavor == "rest" {
		registerFn = "Register" + typ + "ServiceHTTPServer"
	}
	routeLine := "\t" + routePkg + "." + registerFn + "(srv, " + name + "Service)"
	argAnchor := "// register:grpc-arg ── 新模块服务实参在此行后追加(make register 工具锚点,勿删)"
	paramAnchor := "// register:grpc-param ── 新模块服务形参在此行后追加(make register 工具锚点,勿删)"
	if flavor == "rest" {
		argAnchor = "// register:rest-arg ── 新模块服务实参在此行后追加(make register 工具锚点,勿删)"
		paramAnchor = "// register:rest-param ── 新模块服务形参在此行后追加(make register 工具锚点,勿删)"
	}

	// 仓储行(领域服务)与服务客户端行(BFF)二选一,由 wiring.go 里的锚点决定
	repoAnchor := "// ── register:repo ── 新模块仓储在此行后注册(make register 工具锚点,勿删)"
	clientAnchor := "// ── register:client ── 新模块服务客户端在此行后注册(make register 工具锚点,勿删)"
	repoLine := "\t" + name + "Repo := data.New" + typ + "Repo(ctx, entClient)"
	repoSkip := "data.New" + typ + "Repo("
	clientLine := "\t" + name + "ServiceClient := data.New" + typ + "ServiceClient(ctx, discovery)"
	clientSkip := "data.New" + typ + "ServiceClient("

	if err := apply(
		// 仓储(服务客户端)构造行:注入到仓储(客户端)小节末尾
		patch{
			path: wiringPath,
			anchors: []string{
				repoAnchor, repoLine, clientAnchor, clientLine,
			},
			skipIf: repoSkip + "|" + clientSkip,
		},
		// 服务构造行:注入到服务层小节末尾
		patch{
			path: wiringPath,
			anchors: []string{
				"// ── register:service ── 新模块服务在此行后注册(make register 工具锚点,勿删)",
				serviceLine(name, typ, hasAnchor(wiringPath, repoAnchor)),
			},
			skipIf: "service.New" + typ + "Service(",
		},
		// 服务实参(NewGrpcServer / NewRestServer 调用内)
		patch{
			path: wiringPath,
			anchors: []string{
				argAnchor,
				"\t\t" + name + "Service,",
			},
			skipIf: "\t\t" + name + "Service,",
		},
		// server 文件形参
		patch{
			path: serverPath,
			anchors: []string{
				paramAnchor,
				"\t" + name + "Service *service." + typ + "Service,",
			},
			skipIf: name + "Service *service." + typ + "Service,",
		},
		// server 文件路由注册
		patch{
			path: serverPath,
			anchors: []string{
				routeAnchor(flavor),
				routeLine,
			},
			skipIf: registerFn + "(srv, " + name + "Service)",
		},
	); err != nil {
		fmt.Fprintln(os.Stderr, "register:", err)
		os.Exit(1)
	}

	fmt.Printf("已登记 %s(%s):\n"+
		"  %s (repo|client / service / 实参)\n"+
		"  %s (param / route)\n"+
		"下一步: 实现 internal/data/%s_repo.go 与 internal/service/%s_service.go,然后 go build。\n",
		typ, *svc, wiringPath, serverPath, lowerName(typ), lowerName(typ))
}

// serviceLine 生成服务构造行:领域服务依赖仓储,BFF 依赖服务客户端。
// 两种形态的构造在 wiring.go 中各自小节锚点处注入,这里按锚点存在性判断。
func serviceLine(name, typ string, hasRepo bool) string {
	if hasRepo {
		return "\t" + name + "Service := service.New" + typ + "Service(ctx, " + name + "Repo)"
	}
	return "\t" + name + "Service := service.New" + typ + "Service(ctx, " + name + "ServiceClient)"
}

func routeAnchor(flavor string) string {
	if flavor == "rest" {
		return "// register:rest-route ── 新模块服务路由在此行后注册(make register 工具锚点,勿删)"
	}
	return "// register:grpc-route ── 新模块服务路由在此行后注册(make register 工具锚点,勿删)"
}

// findServerFile 在 internal/server 下找含 register:param 锚点的服务端文件,
// 并据锚点判断注册风格(grpc / rest)。
func findServerFile(base string) (path, flavor string, err error) {
	dir := filepath.Join(base, "internal", "server")
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", "", err
	}
	var grpcHits, restHits []string
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		content, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			return "", "", err
		}
		s := string(content)
		if strings.Contains(s, "// register:grpc-param") {
			grpcHits = append(grpcHits, filepath.Join(dir, name))
		}
		if strings.Contains(s, "// register:rest-param") {
			restHits = append(restHits, filepath.Join(dir, name))
		}
	}
	hits := grpcHits
	flavor = "grpc"
	if len(hits) == 0 {
		hits, flavor = restHits, "rest"
	}
	if len(hits) == 0 {
		return "", "", fmt.Errorf("在 %s 下未找到 register:param 锚点文件", dir)
	}
	if len(hits) > 1 {
		fmt.Fprintf(os.Stderr, "register: 多个候选文件 %v,取第一个\n", hits)
	}
	return hits[0], flavor, nil
}

// detectRoutePkg 从服务端文件既有 Register 调用探测路由包别名。
func detectRoutePkg(serverPath string) (string, error) {
	raw, err := os.ReadFile(serverPath)
	if err != nil {
		return "", err
	}
	m := reRouteAlias.FindStringSubmatch(string(raw))
	if m == nil {
		return "", fmt.Errorf("%s: 未探测到路由包别名,请用 -pkg 指定", serverPath)
	}
	return m[1], nil
}

func hasAnchor(path, anchor string) bool {
	raw, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	return strings.Contains(string(raw), anchor)
}

// apply 依次应用各 patch;同文件多个 patch 顺序生效,幂等由 skipIf 保证。
func apply(patches ...patch) error {
	for _, p := range patches {
		if err := applyOne(p); err != nil {
			return err
		}
	}
	return nil
}

func applyOne(p patch) error {
	raw, err := os.ReadFile(p.path)
	if err != nil {
		return err
	}
	content := string(raw)
	for _, skip := range strings.Split(p.skipIf, "|") {
		if skip != "" && strings.Contains(content, skip) {
			return nil // 已登记,幂等跳过
		}
	}

	eol := "\n"
	if strings.Contains(content, "\r\n") {
		eol = "\r\n"
	}
	lines := strings.Split(content, eol)

	for i := 0; i+1 < len(p.anchors); i += 2 {
		anchor, line := p.anchors[i], p.anchors[i+1]
		at := indexOfAnchor(lines, anchor)
		if at < 0 {
			continue
		}
		out := make([]string, 0, len(lines)+1)
		out = append(out, lines[:at+1]...)
		out = append(out, line)
		out = append(out, lines[at+1:]...)
		return os.WriteFile(p.path, []byte(strings.Join(out, eol)), 0o644)
	}
	return fmt.Errorf("%s: 未找到注册锚点 %q,请检查锚点注释是否被移动或删除", p.path, p.anchors[0])
}

func indexOfAnchor(lines []string, anchor string) int {
	for i, line := range lines {
		if strings.TrimSpace(line) == anchor {
			return i
		}
	}
	return -1
}

// pascal 把 sprint / dict_entry / dictEntry 统一为 PascalCase 的实体名。
func pascal(s string) string {
	parts := strings.FieldsFunc(s, func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) })
	var b strings.Builder
	for _, p := range parts {
		r := []rune(p)
		b.WriteRune(unicode.ToUpper(r[0]))
		b.WriteString(string(r[1:]))
	}
	out := b.String()
	if out == "" {
		return s
	}
	return out
}

func lowerName(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	return string(unicode.ToLower(r[0])) + string(r[1:])
}
