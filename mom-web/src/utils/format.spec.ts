import { describe, it, expect } from 'vitest'
import { formatDate, formatNumber, debounce } from './format'

describe('formatDate', () => {
  it('should format date with default format YYYY-MM-DD', () => {
    const date = new Date('2024-03-15T10:30:00')
    expect(formatDate(date)).toBe('2024-03-15')
  })

  it('should format date with custom format', () => {
    const date = new Date('2024-03-15T10:30:00')
    expect(formatDate(date, 'YYYY/MM/DD')).toBe('2024/03/15')
  })

  it('should handle string date input', () => {
    expect(formatDate('2024-03-15')).toBe('2024-03-15')
  })

  it('should handle timestamp input', () => {
    const timestamp = new Date('2024-03-15').getTime()
    expect(formatDate(timestamp)).toBe('2024-03-15')
  })
})

describe('formatNumber', () => {
  it('should format number with 2 decimal places by default', () => {
    expect(formatNumber(3.14159)).toBe('3.14')
  })

  it('should format number with specified decimal places', () => {
    expect(formatNumber(3.14159, 4)).toBe('3.1416')
  })

  it('should handle integer', () => {
    expect(formatNumber(100)).toBe('100.00')
  })
})

describe('debounce', () => {
  it('should delay function execution', async () => {
    let callCount = 0
    const fn = () => { callCount++ }
    const debouncedFn = debounce(fn, 100)

    debouncedFn()
    debouncedFn()
    debouncedFn()

    expect(callCount).toBe(0)

    await new Promise(resolve => setTimeout(resolve, 150))
    expect(callCount).toBe(1)
  })
})
