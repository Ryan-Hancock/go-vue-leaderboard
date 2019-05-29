<template>
  <div class="leaderboard">
      <h1 class="title">CFU - LEADERBAORD</h1>
      <div v-for="item in items" v-bind:key="item.ID">
        <Banner 
          :pos="1" 
          :color="item.colour" 
          :player="item.name"
          :reason="item.last_reason">
        </Banner>
      </div>
  </div>
</template>

<script>
import * as _ from 'lodash'
import axios from 'axios'
import Banner from './../components/Banner.vue'

export default {
  name: 'home',
  data () {
    return {
      items: []
    }
  },
  created: function () {
    this.loadData()
    setInterval(function () {
      this.loadData()
    }.bind(this), 60000)
  },
  methods: {
    loadData: function () {
      axios.get('http://localhost:4505/api/scores')
      .then(response => _.orderBy(response.data.data, 'points', 'desc'))
      .then(res => (this.items = res))
    }
  },
  components: {
    Banner
  }
}
</script>

<style scoped>
  @import url('https://fonts.googleapis.com/css?family=Anton&display=swap');
  
  .leaderboard * {
    font-family: 'Anton', sans-serif !important
  }
  
  .title {
    font-size: 40pt
  }
</style>
