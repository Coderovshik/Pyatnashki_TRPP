<template>
  <div @load="getAuthToken()" class="wrapper">
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
      console.log("Started");
      const data = { email: "test@email.com", password: "qwerty123" };
      let res = await fetch("http://localhost:9999/api/v1/auth/register", {
        method: "POST",
        Headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
      console.log(res.status);
      res = await fetch("http://localhost:9999/api/v1/auth/authenticate", {
        method: "POST",
        Headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
      console.log(res.status);
      let token = await res.json();
      res = await fetch("http://localhost:9999/api/v1/auth/authenticate", {
        method: "GET",
        Headers: {
          Authorization: "Bearer " + token.token,
        },
      });
      data = res.json();
      console.log(data);
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
