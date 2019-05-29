<template>
  <div class="card">
      <UserPortion class="user-item" :pos="pos" :color="color" :player="player"></UserPortion>
      <h3 :style="{float: right}">-{{reason}}</h3>
  </div>
</template>

<script>
import UserPortion from './UserPortion.vue'

export default {
  name: 'banner',
  props: ['pos', 'color', 'player', 'reason'],
  components: {
    UserPortion
  }
}
</script>

<style scoped>
  .card {
  transition: 0.3s;
  width: 100%;
  display: flex;
  flex-direction: row;
  }
  
  .user-item {
    flex-basis: 80%;
  }
</style>
