<template>
  <div class="playing-wrapper">
    <div class="game-area">
      <Plate v-for="num in plates" :key="num" :value="num" />
    </div>
    <button @click="init()" class="game-button">Новая игра</button>
  </div>
</template>

<script>
import Plate from "./Plate.vue";

export default {
  components: { Plate },
  data() {
    return {
      plates: [...Array(16).keys()].map((i) => i + 1),
      timer: 0,
    };
  },
  methods: {
    createMatrix() {
      let x = 0;
      let y = 0;
      for (let i = 0; i < this.plates.length; i++) {
        if (x >= 4) {
          x = 0;
          y++;
        }
        this.$globalMatrix[y][x] = this.plates[i];
        x++;
      }
    },
    setInvisiblePlate() {
      const platesList = Array.from(
        document.querySelectorAll(".plate-wrapper")
      );
      for (let i = 0; i < platesList.length; i++) {
        if (platesList[i].innerText === "16") {
          platesList[i].classList.add("sixteen-plate");
        }
      }
    },
    init() {
      this.plates = this.plates.sort(() => Math.random() - 0.5);
      this.createMatrix();
      setInterval(() => {
        this.$globalTime.value++;
        console.log(this.$globalTime);
      }, 1000);
    },
  },
  mounted() {
    this.setInvisiblePlate();
  },
};
</script>

<style lang="scss" scoped>
.playing-wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.game-area {
  display: grid;
  grid-template-rows: repeat(4, 1fr);
  grid-template-columns: repeat(4, 1fr);
  grid-gap: 10px;
  width: 490px;
  height: 490px;
  background-color: #c6e5a2;
  border: 2px solid #016a70;
  border-radius: 15px;
  box-shadow: 0 0 8px 0 rgba(0, 0, 0, 0.3);
  padding: 10px;
}
.game-button {
  margin-top: 30px;
  border: 2px solid #016a70;
  border-radius: 10px;
  width: 228px;
  height: 60px;
  box-shadow: 0 0 8px 0 rgba(0, 0, 0, 0.3);
  background: #c6e5a2;
  font-family: Gilroy;
  font-weight: 700;
  font-size: 24px;
  line-height: 100%;
  color: #016a70;
  cursor: pointer;
}

.sixteen-plate {
  opacity: 0;
}
</style>
