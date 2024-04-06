<template>
  <div @click="movePlate($event)" class="plate-wrapper">{{ value }}</div>
</template>

<script>
export default {
  props: { value: Number },
  data() {
    return {};
  },
  methods: {
    createMatrix() {
      const platesList = Array.from(
        document.querySelectorAll(".plate-wrapper")
      );
      const platesMatrix = [[], [], [], []];
      let x = 0;
      let y = 0;
      for (let i = 0; i < platesList.length; i++) {
        if (x >= 4) {
          x = 0;
          y++;
        }
        platesMatrix[y][x] = Number(platesList[i].innerText);
        x++;
      }
      return platesMatrix;
    },
    checkSixTeenNear(matrix, plate) {
      let x = 0;
      let y = 0;
      for (let i = 0; i < matrix.length * 4; i++) {
        if (x >= 4) {
          x = 0;
          y++;
        }
        console.log(y, x, matrix[y][x], plate.innerText);
        if (matrix[y][x] === Number(plate.innerText)) {
          const plateCoords = { y, x };
          break;
        }
        x++;
      }
    },
    movePlate(event) {
      const plate = event.target;
      if (plate.innerText === "16") return;
      const matrix = this.createMatrix();
      console.log(this.checkSixTeenNear(matrix, plate));
    },
  },
};
</script>

<style lang="scss">
.plate-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  border: 2px solid #016a70;
  border-radius: 10px;
  width: 110px;
  height: 110px;
  background-color: #d2de32;
  font-family: Gilroy;
  font-weight: 700;
  font-size: 40px;
  color: #016a70;
  cursor: pointer;
  transition: background-color 0.2s;
  &:hover {
    background-color: #f4ff5e;
    box-shadow: 0px 0px 4px #434343;
  }
}
</style>
