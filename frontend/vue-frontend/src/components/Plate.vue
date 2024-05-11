<template>
  <div @click="movePlates($event)" class="plate-wrapper">
    {{ value }}
  </div>
</template>

<script>
import AXIOS from "@/http-common";

export default {
  props: { value: Number },
  data() {
    return {};
  },
  methods: {
    checkWin(matrix) {
      let x = 0;
      let y = 0;
      let win = true;
      for (let i = 0; i < matrix.length * 4; i++) {
        if (x >= 4) {
          x = 0;
          y++;
        }
        if (matrix[y][x] !== i + 1) {
          win = false;
        }
        x++;
      }
      if (win) {
        clearInterval(this.$globalTimerID.value);
        const plates = document.querySelectorAll(".plate-wrapper");
        for (let i = 0; i < plates.length; i++) {
          plates[i].style.display = "none";
        }
        document.querySelector(".game-area").style.display = "flex";
        document.querySelector(".win-wrapper").style.display = "flex";
        const score = { score: this.$globalMoves + this.$globalTime };
        AXIOS.post(
          "http://localhost:8081/scores?user_id=" +
            localStorage.getItem("UserId"),
          score,
          {
            headers: {
              "Content-type": "application/json",
            },
          }
        );
      }
    },
    findCoods(matrix, plate) {
      let x = 0;
      let y = 0;
      let plateCoords = {};
      let sixteenCoords = {};
      for (let i = 0; i < matrix.length * 4; i++) {
        if (x >= 4) {
          x = 0;
          y++;
        }
        if (matrix[y][x] === Number(plate.innerText)) {
          plateCoords = { y, x };
        } else if (matrix[y][x] === 16) {
          sixteenCoords = { y, x };
        }
        x++;
      }
      return { plateCoords, sixteenCoords };
    },
    checkMovable(plate) {
      if (plate.innerText === "16") return false;
      const coords = this.findCoods(this.$globalMatrix, plate);
      let result = false;
      if (
        (coords.plateCoords.y === coords.sixteenCoords.y &&
          Math.abs(coords.plateCoords.x - coords.sixteenCoords.x) === 1) ||
        (coords.plateCoords.x === coords.sixteenCoords.x &&
          Math.abs(coords.sixteenCoords.y - coords.plateCoords.y) === 1)
      ) {
        result = true;
        return { result, coords };
      } else {
        return result;
      }
    },
    movePlates(event) {
      const plate = event.target;
      const movable = this.checkMovable(plate);
      const result = movable.result;
      const coords = movable.coords;
      if (!result) return;
      const plateCoords = coords.plateCoords;
      const sixteenCoords = coords.sixteenCoords;
      const sixTeenPlate = document.querySelector(".sixteen-plate");
      plate.style.gridColumn = sixteenCoords.x + 1;
      plate.style.gridRow = sixteenCoords.y + 1;
      sixTeenPlate.style.gridColumn = plateCoords.x + 1;
      sixTeenPlate.style.gridRow = plateCoords.y + 1;
      this.$globalMatrix[sixteenCoords.y][sixteenCoords.x] = Number(
        plate.innerText
      );
      this.$globalMatrix[plateCoords.y][plateCoords.x] = 16;
      this.$globalMoves.value++;
      this.checkWin(this.$globalMatrix);
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
