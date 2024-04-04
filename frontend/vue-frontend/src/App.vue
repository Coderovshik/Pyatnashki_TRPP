<template>
  <div class="wrapper">
    <Header />
    <Game />
  </div>
</template>

<script>
import Header from "./components/Header.vue";
import Game from "./components/Game.vue";

export default {
  name: "App",
  components: { Header, Game },
};
</script>

<style lang="scss">
#app {
  display: flex;
  flex-direction: column;
  min-width: 100vw;
  min-height: 100vh;
  background-color: #ffffdd;
}
</style>
