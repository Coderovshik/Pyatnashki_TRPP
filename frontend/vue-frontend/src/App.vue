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
  data() {
    return {
      SECRET_KEY:
        "92e17dc8bfac0842d5068f383dd93181e78ee27d3abdb3ed6df848c134b446e5",
    };
  },
  methods: {
    async getAuthToken() {
      const res = await fetch("http://localhost:9999/api/v1/auth/register");
      const data = await res.json();
      const token = data.token;
    },
  },
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
