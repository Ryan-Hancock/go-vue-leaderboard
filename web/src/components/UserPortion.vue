<template>
  <div
    class="userport"
  >
  <div class="user-content">
    <h1 v-bind:style="{ color: color}">#{{pos}} {{player}}</h1>
  </div>
    
  </div>
</template>

<script>
export default {
  props: ['pos', 'color', 'player']
}
</script>

<style scoped>
.userport {
  
}

.user-content, * {
  float: left;
  font-size: 40pt;
}
</style>
