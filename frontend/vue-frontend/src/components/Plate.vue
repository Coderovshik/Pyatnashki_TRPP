<template>
  <div class="plate-wrapper">{{ value }}</div>
</template>

<script>
export default {
  props: { value: Number },
  data() {
    return {};
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
}
</style>
