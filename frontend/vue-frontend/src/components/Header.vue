<template>
  <header class="header">
    <h1 class="game-name">Tag Game</h1>
    <div class="user">
      <h2>{{ username }}</h2>
      <div class="user-icon">
        <fa icon="user" class="fa-user" />
      </div>
      <h2 class="user-name"></h2>
    </div>
  </header>
</template>

<script>
export default {
  data() {
    return {
      username: "Artem",
    };
  },
};
</script>

<style scoped lang="scss">
$dark-green: #016a70;
* {
  font-family: Gilroy;
}
.header {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  margin: 0 25px;
  padding-top: 20px;
  .game-name {
    font-size: 48px;
    font-weight: bold;
    color: $dark-green;
    margin: 0;
  }
  .fa-user {
    font-size: 28px;
    color: $dark-green;
  }
  .user-icon {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 52px;
    height: 52px;
    border-radius: 50%;
    border: 2px solid $dark-green;
    margin-left: 10px;
  }
  .user {
    display: flex;
    flex-direction: row;
    align-items: center;
    h2 {
      color: $dark-green;
      font-size: 24px;
      font-weight: bold;
    }
  }
}
</style>
