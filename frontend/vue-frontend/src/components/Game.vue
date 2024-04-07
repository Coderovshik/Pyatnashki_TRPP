<template class="game">
  <div class="game-wrapper">
    <Leaderboard />
    <PlayingArea />
    <Score :timer="$globalTime" :moves="$globalMoves" />
  </div>
</template>

<script>
import Leaderboard from "./Leaderboard.vue";
import PlayingArea from "./PlayingArea.vue";
import Score from "./Score.vue";

export default {
  components: { Leaderboard, PlayingArea, Score },
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
  justify-content: space-around;
}
</style>
