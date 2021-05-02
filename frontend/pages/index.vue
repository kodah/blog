<template>
  <section class="section">
    <div class="column is-rounded has-background-black-ter is-offset-2 is-8">
      <article class="columns featured">
        <div class="column is-12 featured-content">
          <div>
            <h3 class="heading has-text-primary">Category</h3>
            <h1 class="title has-text-primary-light">Title</h1>
            <p class="post-excerpt has-text-primary-light ">Summary</p>
            <br>
            <a href="#" class="button is-primary has-text-primary-light">Read More</a>
          </div>
        </div>
      </article>
    </div>
  </section>
</template>

<script>
import Card from '~/components/Card'

export default {
  name: 'HomePage',
  layout: 'default',
  components: {
    Card,
  },
  auth: 'guest',
}
</script>
