import { createApp } from "vue";
import App from "./App.vue";

import { library } from "@fortawesome/fontawesome-svg-core";
import { fas } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { ref } from "vue";

library.add(fas);
let app = createApp(App);
app.config.globalProperties.$globalMatrix = [[], [], [], []];
app.config.globalProperties.$globalTime = ref(0);
app.config.globalProperties.$globalMoves = ref(0);
app.config.globalProperties.$globalTimerID = ref(0);
app.component("fa", FontAwesomeIcon);
app.mount("#app");

async function getAuthToken() {
  console.log("Started");
  const data = { email: "test@email.com", password: "qwerty123" };
  console.log(JSON.stringify(data));
  let res = await fetch("http://localhost:9999/api/v1/auth/register", {
    method: "POST",
    mode: "cors",
    cache: "no-cache",

    Headers: {
      "Content-Type": "application/json; charset=utf-8",
    },
    body: JSON.stringify(data),
  });
  console.log(res);
  res = await fetch("http://localhost:9999/api/v1/auth/authenticate", {
    method: "POST",
    Headers: {
      "Content-Type": "application/json; charset=utf-8",
    },
    body: JSON.stringify(data),
  });
  console.log(res);
  let token = await res.json();
  res = await fetch("http://localhost:9999/api/v1/auth/authenticate", {
    method: "GET",
    Headers: {
      Authorization: "Bearer " + token.token,
    },
  });
  data = res.json();
  console.log(data);
}

//document.addEventListener("DOMContentLoaded", getAuthToken);
