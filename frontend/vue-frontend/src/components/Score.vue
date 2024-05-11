<template>
  <div class="score-wrapper">
    <h2>Текущая статистика</h2>
    <div class="score">
      <div class="time">
        <h2 class="digits">{{ timer }}</h2>
        <h2>секунд</h2>
      </div>
      <div class="moves">
        <h2 class="digits">{{ moves }}</h2>
        <h2>ходов</h2>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: { timer: Object, moves: Object },
  data() {
    return {};
  },
  methods: {},
};
</script>

<style scoped lang="scss">
.score-wrapper {
  display: flex;
  flex-direction: column;
  box-shadow: 0 0 8px 0 rgba(0, 0, 0, 0.3);
  background: #c6e5a2;
  border: 2px solid #016a70;
  border-radius: 15px;
  width: 280px;
  height: 231px;
  h2 {
    text-align: center;
    font-family: Gilroy;
    color: #016a70;
    margin-top: 10px;
  }
  .digits {
    font-size: 40px;
    margin-bottom: 10px;
  }
  .score {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-around;
  }
}
</style>
