/**
 * 分析结果 CSV 导出（纯前端，数据已在手，无需后端参与）
 *
 * RFC4180：逗号/引号/换行转义，UTF-8 BOM 保证 Excel 打开不乱码。
 */

export function exportCsv(filename: string, headers: string[], rows: (string | number | null | undefined)[][]): void {
  const csv = [
    headers.map(quoteCell).join(','),
    ...rows.map((row) => row.map(quoteCell).join(',')),
  ].join('\r\n');

  // \uFEFF BOM：让 Excel 以 UTF-8 解析中文
  const blob = new Blob([`\uFEFF${csv}`], { type: 'text/csv;charset=utf-8;' });
  const url = URL.createObjectURL(blob);

  const link = document.createElement('a');
  link.href = url;
  link.download = `${filename}_${new Date().toISOString().slice(0, 10)}.csv`;
  link.style.display = 'none';
  document.body.append(link);
  link.click();
  link.remove();
  URL.revokeObjectURL(url);
}

function quoteCell(value: string | number | null | undefined): string {
  const str = value === null || value === undefined ? '' : String(value);
  const escaped = str.replaceAll('"', '""').replaceAll('\n', ' ');
  return `"${escaped}"`;
}
