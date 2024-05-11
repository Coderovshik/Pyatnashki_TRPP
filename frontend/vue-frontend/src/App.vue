<template>
  <div class="wrapper">
    <RegWindow />
    <Header />
    <Game />
  </div>
</template>

<script>
import RegWindow from "./components/RegWindow.vue";
import Header from "./components/Header.vue";
import Game from "./components/Game.vue";

export default {
  name: "App",
  components: { RegWindow, Header, Game },
  data() {
    return {};
  },
};
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
</style>
