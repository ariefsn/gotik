import { describe, expect, it } from 'vitest'

import { mount } from '@vue/test-utils'
import Component from '../H1.vue'

describe('H1', () => {
  it('renders properly', () => {
    const wrapper = mount(Component, { slots: { default: 'H1' } })
    expect(wrapper.text()).toContain('H1')
  })
})
