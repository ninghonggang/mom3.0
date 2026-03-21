import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import { defineComponent, h } from 'vue'

// 创建一个简单的 HelloWorld 组件用于测试
const HelloWorld = defineComponent({
  name: 'HelloWorld',
  props: {
    msg: {
      type: String,
      default: 'Hello'
    },
    count: {
      type: Number,
      default: 0
    }
  },
  setup(props) {
    return () => h('div', { class: 'hello-world' }, [
      h('h1', props.msg),
      h('p', `Count: ${props.count}`)
    ])
  }
})

describe('HelloWorld Component', () => {
  it('renders properly with default props', () => {
    const wrapper = mount(HelloWorld)
    expect(wrapper.find('h1').text()).toBe('Hello')
    expect(wrapper.find('p').text()).toBe('Count: 0')
  })

  it('renders with custom msg prop', () => {
    const wrapper = mount(HelloWorld, {
      props: { msg: 'Welcome', count: 5 }
    })
    expect(wrapper.find('h1').text()).toBe('Welcome')
    expect(wrapper.find('p').text()).toBe('Count: 5')
  })

  it('updates when props change', async () => {
    const wrapper = mount(HelloWorld, {
      props: { msg: 'Start', count: 0 }
    })

    await wrapper.setProps({ count: 10 })
    expect(wrapper.find('p').text()).toBe('Count: 10')

    await wrapper.setProps({ msg: 'Updated' })
    expect(wrapper.find('h1').text()).toBe('Updated')
  })

  it('has correct CSS class', () => {
    const wrapper = mount(HelloWorld)
    expect(wrapper.find('.hello-world').exists()).toBe(true)
  })
})
