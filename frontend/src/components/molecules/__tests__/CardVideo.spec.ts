import { describe, expect, it } from 'vitest'

import { mount } from '@vue/test-utils'
import Component from '../CardVideo.vue'

describe('CardVideo', () => {
  it('renders properly', () => {
    const wrapper = mount(Component, {
      props: {
        video: {
          common: {
            doc_id_str: '1234567890',
          }
        }
      }
    })
    expect(wrapper.isVisible()).toBe(true)
  })
})
