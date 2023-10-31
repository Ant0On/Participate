<template>
  <button :class="['button', className]">
    <label_component class="label-instance" :label-class-name="labelLabelClassName" :text="labelText" />
  </button>
</template>

<script>
import label_component from "./Label.vue";

export default {
  name: "button_component",
  components: {
    label_component,
  },
  props: {
    className: {
      type: String,
      default: "",
    },
    labelText: {
      type: String,
      default: "Button",
    },
    labelLabelClassName: {
      type: String,
      default: "",
    },
  },
};
</script>

<style>
.button {
  all: unset;
  align-items: center;
  background-color: var(--text-link);
  border-radius: 6px;
  box-sizing: border-box;
  display: inline-flex;
  flex-direction: column;
  justify-content: center;
  overflow: hidden;
  padding: 4px 16px;
  position: relative;
}

.button .label-instance {
  flex: 0 0 auto !important;
}
</style>