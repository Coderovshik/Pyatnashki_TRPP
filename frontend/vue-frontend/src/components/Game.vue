<template class="game">
  <div class="game-wrapper">
    <PlayingArea />
  </div>
</template>

<script>
import PlayingArea from "./PlayingArea.vue";

export default {
  components: { PlayingArea },
  data() {
    return {};
  },
};
</script>

<style lang="scss" scoped>
.game-wrapper {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}
</style>
