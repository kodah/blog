<template>
  <div>
    <b-navbar type="is-dark">
      <template #brand>
        <b-navbar-item class="is-lowercase pl-4 pr-4 has-text-primary-light has-text-weight-bold is-active has-background-primary" tag="router-link" :to="{ path: '/' }">
          kodah's blog
        </b-navbar-item>
      </template>

      <template #start>
            <nuxt-link v-for="(item, key) of items" :key="key" class="navbar-item" :to="item.to">
              <b-icon class="has-text-white" :icon="item.icon" /> <span class="has-text-primary-light">{{ item.title }}</span>
            </nuxt-link>
      </template>

      <template #end>
        <b-navbar-item tag="div">
          <AccountDropdown />
        </b-navbar-item>
      </template>
    </b-navbar>
    <nuxt />
  </div>
</template>

<script>
import AccountDropdown from "~/components/AccountDropdown";

export default {
  name: "default",
  components: {
    AccountDropdown
  },
  data() {
    return {
      items: [
        {
          title: 'Home',
          icon: 'home',
          to: { name: 'index' },
        },
        {
          title: 'About',
          icon: 'lightbulb',
          to: { name: 'about' },
        },
      ],
    }
  },
}
</script>

<style lang="scss">
@import '~assets/scss/variables.scss';

:root {
  background: $background;
}
</style>
