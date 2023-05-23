import Home from "../../src/views/Home.vue";
import { createPinia } from "pinia";

describe('Home.cy.js', () => {
  it('playground', () => {
    cy.mount(Home, {
      global: {
        plugins: [createPinia()]
      }
    });
  })
})