<template>
  <div :class="['nav-item', className]" @mouseover="handleMouseOver" @mouseout="handleMouseOut">
    <div class="home">{{ text }}</div>
  </div>
</template>

<script>
export default {
  name: "nav_item_component",
  props: {
    className: {
      type: String,
      default: "",
    },
    text: {
      type: String,
      default: "Home",
    },
  },
  methods: {
    handleMouseOver() {
      this.$el.style.backgroundColor = "#3fc1c9";
    },
    handleMouseOut() {
      this.$el.style.backgroundColor = "#efefef";
    },
  },
};
</script>

<style>
.nav-item {
  align-items: center;
  background-color: #efefef;
  border-radius: 5px;
  display: inline-flex;
  gap: 10px;
  justify-content: center;
  padding: 12px 26px;
  position: relative;
}

.nav-item .home {
  color: #000000;
  font-family: "Poppins", Helvetica;
  font-size: 20px;
  font-weight: 400;
  letter-spacing: 0;
  line-height: normal;
  margin-top: -1px;
  position: relative;
  width: fit-content;
}
</style>
