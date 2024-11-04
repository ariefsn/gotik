import { describe, expect, it } from 'vitest'

import { mount } from '@vue/test-utils'
import Component from '../Button.vue'

describe('Button', () => {
  it('renders properly', () => {
    const wrapper = mount(Component, { slots: { default: 'Button' } })
    expect(wrapper.text()).toContain('Button')
  })
})
