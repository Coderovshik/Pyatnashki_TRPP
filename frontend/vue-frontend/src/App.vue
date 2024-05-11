<template>
  <div class="wrapper">
    <!-- <RegWindow /> -->
    <Header />
    <Game />
  </div>
</template>

<script>
import Header from "./components/Header.vue";
import Game from "./components/Game.vue";
import AXIOS from "@/http-common";

export default {
  name: "App",
  components: { Header, Game },
  data() {
    return {};
  },
  methods: {
    async getAuthToken() {
      const data = { email: "test@email.com", password: "qwerty123" };
      AXIOS.post("http://localhost:9999/api/v1/auth/register", data, {
        Headers: {
          "Content-type": "application/json;charset=utf-8",
          Accept: "application/json, text/plain",
        },
      });
      console.log("Started");

      // let res = await fetch("http://localhost:9999/api/v1/auth/register", {
      //   method: "POST",
      //   mode: "cors",
      //   Headers: {
      //     "Content-Type": "application/json",
      //   },
      //   body: JSON.stringify(data),
      // });
      // console.log(res.status);
      // res = await fetch("http://localhost:9999/api/v1/auth/authenticate", {
      //   method: "POST",
      //   Headers: {
      //     "Content-Type": "application/json",
      //   },
      //   body: JSON.stringify(data),
      // });
      // console.log(res.status);
      // let token = await res.json();
      // res = await fetch("http://localhost:9999/api/v1/auth/authenticate", {
      //   method: "GET",
      //   Headers: {
      //     Authorization: "Bearer " + token.token,
      //   },
      // });
    },
  },
  mounted() {
    this.getAuthToken();
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
