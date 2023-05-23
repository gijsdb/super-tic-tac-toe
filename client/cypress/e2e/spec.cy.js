
const users = [{ user: 'user1', username: "" }, { user: 'user2', username: "" }]

describe('spec.cy.js', () => {
  it('assigns a new user name for different users', () => {
    cy.wrap(users).each((user) => {
      cy.visit("http://localhost:3000/")
      cy.get('#username').then((content) => {
        user.username = content.text()
      })
    }).then(() => {
      const usernames = users.map(user => user.username);
      const uniqueUsernames = [...new Set(usernames)];
      expect(uniqueUsernames.length).to.eq(users.length);
    })
  })
})