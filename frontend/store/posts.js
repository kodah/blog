export const state = () => ({
  list: []
})

export const mutations = {
  add(state, post) {
    state.list.push({
      post,
    })
  },
  remove(state, { todo }) {
    state.list.splice(state.list.indexOf(todo), 1)
  },
  // performs a full fetch of posts
  async fetchPosts(state) {
    state.list = [];

    const posts = await this.$axios.$get("/api/v1/");
  }
}
