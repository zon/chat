export const locale = 'en-US'

const options: Intl.DateTimeFormatOptions = {
  month: 'numeric',
  day: '2-digit',
  year: '2-digit',
  hour: '2-digit',
  minute: '2-digit',
}
const format = new Intl.DateTimeFormat(locale, options)

export function formatDate(date: Date): string {
  return format.format(date)
}
