<template>
  <div v-if="!authorized" class="reg-window">
    <from class="reg-form reg"
      ><h1>Регистрация</h1>
      <h2>Почта</h2>
      <input
        type="email"
        class="email"
        placeholder="tag@game.mail"
        v-model="formData.email"
        required
        pattern="^\S+@\S+\.\S+$"
      />
      <h2>Пароль</h2>
      <input
        type="password"
        class="password"
        placeholder="yourpassword"
        v-model="formData.password"
        required
        pattern="^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,30}$"
      />
      <h2>Повторите пароль</h2>
      <input
        type="password"
        class="password password-again"
        placeholder="yourpassword"
        v-model="passwordAgain"
        required
        pattern="^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,30}$"
      />
      <button @click.prevent="register()" class="register-btn">
        Регистрация
      </button>
    </from>
    <from class="log-form reg-form"
      ><h1>Войти</h1>
      <h2>Почта</h2>
      <input
        type="email"
        class="email"
        placeholder="tag@game.mail"
        v-model="formData.email"
        required
      />
      <h2>Пароль</h2>
      <input
        type="password"
        class="password"
        placeholder="yourpassword"
        v-model="formData.password"
        required
      />
      <button @click.prevent="login()" class="register-btn">Войти</button>
    </from>
  </div>
</template>

<script>
import AXIOS from "@/http-common";

export default {
  data() {
    return {
      formData: {
        email: "",
        password: "",
      },
      passwordAgain: "",
      authorized: localStorage.getItem("authorized"),
    };
  },
  methods: {
    checkInputs() {
      const requiredInputs = Array.from(document.querySelectorAll(":required"));
      for (const input of requiredInputs) {
        if (input.checkValidity() === false) return false;
      }
      return true;
    },
    closeForm(className) {
      document.querySelector(className).style.display = "none";
    },
    openLogForm() {
      document.querySelector(".log-form").style.display = "flex";
    },
    register() {
      if (this.checkInputs()) {
        AXIOS.post(
          "http://localhost:9999/api/v1/auth/register",
          this.formData,
          {
            headers: {
              "Content-type": "application/json",
            },
          }
        )
          .then(() => {
            alert("Вы успешно зарегестрировалась!");
            this.closeForm(".reg");
            this.openLogForm();
          })
          .catch((e) => {});
      } else alert("Ошибка! Неверные данные.");
    },
    login() {
      AXIOS.post(
        "http://localhost:9999/api/v1/auth/authenticate",
        this.formData,
        {
          headers: {
            "Content-type": "application/json",
          },
        }
      )
        .then((response) => {
          localStorage.setItem("authorized", true);
          localStorage.setItem("userId", response.data.id);
          localStorage.setItem("userName", response.data.username);
          this.closeForm(".log-form");
          document.querySelector(".reg-window").style.display = "none";
        })
        .catch((e) => {});
    },
  },
};
</script>

<style lang="scss" scoped>
.reg-window {
  display: flex;
  align-items: center;
  justify-content: center;
  position: absolute;
  z-index: 1;
  width: 100vw;
  height: 100vh;
  backdrop-filter: blur(6px);
}
.reg-form {
  display: flex;
  flex-direction: column;
  align-items: center;
  box-shadow: 0 4px 4px 0 rgba(0, 0, 0, 0.25);
  background: #c6e5a2;
  border: 2px solid #016a70;
  border-radius: 15px;
  width: 450px;
  height: 526px;
  font-family: Gilroy;
  h1 {
    font-weight: 700;
    font-size: 36px;
    color: #016a70;
  }
  h2 {
    font-weight: 600;
    font-size: 22px;
    color: #016a70;
    text-align: start;
    width: 363px;
    margin: 15px auto;
  }
  input {
    font-family: Gilroy;
    border: 2px solid #ffd;
    border-radius: 10px;
    width: 363px;
    height: 45px;
    background: #ffd;
    font-size: 18px;
    outline: none;
    transition: 0.1s;
    padding-inline-start: 15px;
    &::placeholder {
      font-weight: 600;
      font-size: 18px;
      color: rgba(0, 0, 0, 0.29);
    }
    &:focus {
      border: 2px solid #016a70;
    }
  }
  .register-btn {
    border: 2px solid #016a70;
    border-radius: 10px;
    width: 383px;
    height: 50px;
    background: #d2de32;
    font-weight: 600;
    font-size: 21px;
    color: #016a70;
    margin-top: 20px;
    cursor: pointer;
    transition: 0.2s;
    &:hover {
      box-shadow: 0 3px 4px 0 rgba(0, 0, 0, 0.25);
    }
  }
}
.log-form {
  display: none;
}
</style>
