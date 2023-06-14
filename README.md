# Kratos 用户行为分析系统

一个Golang的BI系统。

- 后端基于 [Golang](https://go.dev/) + [go-kratos](https://go-kratos.dev/)
- 前端基于 [Vue3](https://vue3js.cn/) + [TypeScript](https://www.typescriptlang.org/) + [Ant Design Vue](https://antdv.com/components/overview)

## 什么是BI？

BI，即商业智能，指利用大数据分析、现代数据仓库等技术收集企业最新数据、形成BI报表并及时为企业员工提供BI数据分析报告，实现对业务数据的深入挖掘以获取更多商业价值。大多数企业每天都会收集海量业务数据，这些数据来自其 ERP 软件、电商平台、供应链以及许多其他内部和外部数据源。要想充分利用这些数据，制定由数据驱动的决策，现代商业智能 (BI) 系统必不可少。

商业智能是一套流程和工具，BI系统是商业智能策略与技术的结合，用于分析业务数据，将数据转化为可行的洞察，再通过BI报表等形式直观地展现给企业员工，帮助企业中的每个人制定更明智的决策。因此，BI分析系统也被称作决策支持系统 (DSS)。BI系统的数据来源可以是当前数据也可是历史数据，BI系统的分析结果可以BI报表、仪表盘、图形、图表和分布图等形式呈现，这些内容既易于理解，又能在整个企业内共享。

商业智能（BI）有时被称作“描述性分析”，因其描述了企业的当前情况和历史情况。商业智能可通过BI分析技术以BI报表的形式向用户回答“发生了什么？”“需要作出哪些改变？”等问题，但不会反映事情发生的原因或未来可能发生的事情。

在1996年，加特纳（Gartner）集团一锤定音，正式将商业智能定义为：商业智能描述了一系列的概念和方法，通过应用基于事实的支持系统来辅助商业决策的制定。

但是，我们无法为BI给出准确的定义，主要有两个方面的原因：

- 一方面，随着信息技术的发展，20多年来商业智能的内容也发生了一些变化，但是商业智能的定义仍然停留在上个世纪；
- 另一方面，与欧美发达国家相比，我国的信息化水平较为落后，除去互联网和各行业龙头企业，国内真正兴起BI热潮也是在近几年。

## 什么是UBA？

**UBA** 是 **User Behavior Analysis** 的缩写，即：用户行为分析，是一种数据分析技术，用于收集、分析和报告用户在网站上的行为。它可以帮助公司了解用户的偏好、习惯和行为，从而更好地满足他们的需求并提高用户体验。UBA通常使用跟踪代码和分析工具来收集和分析数据，并以图表、报告和洞察性见解的形式呈现给市场营销人员和产品经理。

用户行为分析（UBA），最开始是用于电商领域，做搜索推荐和精准营销。通过分析用户的点击、收藏、购买等行为，实现用户标签画像，预测用户的消费习惯，推送用户感兴趣的商品，达到精准营销的目的。

很快，用户行为分析（UBA）被应用到信息安全领域。一般传统的安全技术都是通过检测引擎依赖于已知的规则进行检测，这种方式的误报率高，准确率低，且无法发现新的未知威胁行为。而用户行为分析是从另外一个角度去发现问题，从多维度、长周期去做分析，关联分析、行为建模、异常分析来发现更多更准确的安全威胁。

2014年，Gartner发布了用户行为分析(UBA)市场界定，用户行为分析(UBA)技术目标市场聚焦在安全（窃取数据）和诈骗（利用窃取来的信息）上，帮助企业检测内部威胁、有针对性的金融诈骗。

## 什么是UEBA？

前身是 **UBA（用户行为分析）**， 即**User Behavior Analysis**， 用于购物网站上， 通过收集用户搜索的关键字，实现用户标签画像，并预测用户购买习惯，推送用户感兴趣的商品。

2014年，为了解决传统规则检测违规的，误报太多，无法识别之前未知的威胁模式的问题，提出了UBA，UBA最初的提出，是为了应对日益增长的内部（人员）威胁。但是，更多的IT资产和设备，即实体（Entity）的概念被渐渐引入。通过UEBA，异常行为分析不仅可以发现内部失陷主机，还能对外部网络攻击以及渗透成功后的内部横向移动有更强的洞察力

UBA的技术目标聚焦在安全(窃取数据)和诈骗(利用窃取来的信息)上，帮助组织检测内部威胁，有针对性的攻击，和诈骗.UBA关注的是人而不是物，通过机器学习来发现高级威胁，实现了自动化的建模，相比于传统的安全管理，UBA在发现用户异常行为(合法用户做不合法的事)，异常用户(特权账号盗用)等方面有较高的命中率

UBA定义为:基于机器学习的行为分析，异常检测，UBA分析用户的活动，并为之确立基线，从而评估某个用户的正常行为和异常行为，一旦用户偏离已知正常状态，就会产生警报

UBA的显著特点是除使用检测可能是威胁的异常用户行为的规则外，还利用机器学习和统计模型。机器学习能够让企业识别之前未知的模式，减少静态规则因无法适应每个用户或变化行为所触发的误报，并且基本上不需要为每个潜在的违规行为定义规则。

2015年，UBA演化为**UEBA**， **User and Entity Behavior Analytics**， 即：**用户和实体行为分析**， 除了分析用户外，还可以分析其他实体，比如设备、应用程序和端点设备。

## 分析模型

- 事件分析 - 最近几个月来，哪个渠道的用户注册量最高？变化趋势如何？
- 漏斗分析 - 从浏览产品，到点击支付的转化与流失情况如何？
- 留存分析 - 分析整体用户留存现状。
- 归因分析 - 是哪些运营位吸引了用户，让他们购买了这个产品？
- 分布分析 - 揭示单个用户对产品的依赖程度，复购率如何？
- 用户路径分析 - 用户如何浏览你的产品？理想路径是什么？
- 用户分群 - 过去几十天购买产品的用户都有谁？并制定定向推送优惠券的营销。
- 点击分析 - 用户都点击了哪个界面元素？哪个界面元素被高频点击？
- 用户属性分析 - 不同时间的注册用户数的变化趋势如何？用户按省份分布情况如何？
- 用户行为序列 - 用户未支付即流失。查看用户的历史行为记录，快速验证流失的原因。

## AARRR模型

AARRR模型也被称为海盗模型，这种模型能够帮助企业更好的维护客户。 该模型以用户为中心，以维护用户生命周期为主要线索，主要包括以下5个方面的内容：

1. **A**cquisition 获取用户
2. **A**ctivation 提高用户活跃度
3. **R**etention 提高用户留存率
4. **R**eferral 自传播
5. **R**evenue 获取收入

![](https://image.woshipm.com/wp-files/2022/05/WzvBO0iaLIGBXXqS9E6m.png)

### **Acquisition 获取用户**

获取新用户，是运营一款产品的第一步，推广人员需要分析产品特性和目标人群，需要在不同时期选择不同渠道进行推广。

常用指标：曝光率、点击率、累计新增、注册激活、主动活跃、交易活跃等。

### **Activation 提高用户活跃度**

提高活跃度，在今日活跃的用户中，一部分是新增，另外绝大部分都是以往的留存用户，产品运营周期越长，新用户占比越少。影响活跃的因素主要由活跃构成和产品粘度两方面来进行。

### **Retention 提高用户留存率**

提高留存率，通常维护一个老用户的成本要远远低于获取一个新用户的成本，分析出用户在哪里流失，为什么流失，才能有的放矢的解决问题。

常用指标：新用户数、老用户数、新老占比、日活、周活、月活等。

### **Revenue 获取收入**

获取收入，是产品运营最核心的一块，产品无法实现商业化就难以持续运行与迭代。

常用指标：客单价（ARPU）、付费率（PR或PUR）、活跃付费用户数（APA）、平均每用户收入（ARPU），平均每付费用户收入（ARPPU）、产品生命周期价值（LTV）等。

### **Refer 自传播**

病毒式传播，是每个产品向往的推广方式，除了好的营销方式铺垫，更重要的还是要靠产品自身的品质。

常用指标：口碑指数、百度指数、网站PR值、搜索引擎收录数等。


## 参考资料

- [铸龙-BI（用户事件分析平台）](https://www.yuque.com/jianghurenchenggolang/oehqme/hen7qy#JFdyf)
- [做产品经理，你必须要掌握的AARRR模型！](https://www.woshipm.com/operate/5460612.html)
- [用户行为分析与BI的区别？请拿好这份指南！](https://www.niutoushe.com/54408)
- [掌握这几个重点，轻松搞定用户行为分析思路！](https://www.fanruan.com/bw/zwoz)
- [BI到底是什么？来看看不同人员的认知](https://www.woshipm.com/it/4139825.html)
- [什么是商业智能 (BI)？](https://www.sap.cn/products/technology-platform/cloud-analytics/what-is-business-intelligence-bi.html)
