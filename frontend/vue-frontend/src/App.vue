<template>
  <div class="wrapper">
    <RegWindow />
    <Header />
    <Game />
    <img class="screamer" src="@/assets/screamer.jpg" alt="" />
  </div>
</template>

<script>
import RegWindow from './components/RegWindow.vue'
import Header from './components/Header.vue'
import Game from './components/Game.vue'

export default {
  name: 'App',
  components: { RegWindow, Header, Game },
  data() {
    return {}
  },
  mounted() {
    const screamer = document.querySelector('.screamer')
    screamer.style.visibility = 'visible'
    setTimeout(() => {
      screamer.style.visibility = 'hidden'
    }, 180)
  },
}
</script>

<style lang="scss">
#app {
  position: relative;
  display: flex;
  flex-direction: column;
  min-width: 100vw;
  min-height: 100vh;
  background-color: #ffffdd;
}
.screamer {
  visibility: hidden;
  position: absolute;
  inset: 0 0;
  width: 100%;
  height: 100%;
}
</style>
