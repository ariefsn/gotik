import { describe, expect, it } from 'vitest'

import { mount } from '@vue/test-utils'
import Component from '../H2.vue'

describe('H2', () => {
  it('renders properly', () => {
    const wrapper = mount(Component, { slots: { default: 'H2' } })
    expect(wrapper.text()).toContain('H2')
  })
})
