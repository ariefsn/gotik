import { describe, expect, it } from 'vitest'

import { mount } from '@vue/test-utils'
import Component from '../InputText.vue'

describe('InputText', () => {
  it('renders properly', () => {
    const wrapper = mount(Component, { props: { modelValue: 'InputText', placeholder: 'Search here...' } })
    expect(wrapper.props('placeholder')).toContain('Search here...')
  })
})
