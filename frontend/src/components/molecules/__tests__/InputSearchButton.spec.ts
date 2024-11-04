import { describe, expect, it } from 'vitest'

import { mount } from '@vue/test-utils'
import Component from '../InputSearchButton.vue'

describe('InputSearchButton', () => {
  it('renders properly', () => {
    const wrapper = mount(Component, { props: { placeholder: 'Search here...', modelValue: 'test' } })
    expect(wrapper.isVisible()).toBe(true)
  })
})
