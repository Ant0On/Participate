<template>
  <div :class="['label', className]">
    <div :class="['text-wrapper-2', labelClassName]">{{ text }}</div>
  </div>
</template>

<script>
export default {
  name: "label_component",
  props: {
    className: {
      type: String,
      default: "",
    },
    labelClassName: {
      type: String,
      default: "",
    },
    text: {
      type: String,
      default: "Button",
    },
  },
};
</script>

<style>
.label {
  align-items: center;
  display: inline-flex;
  position: relative;
}

.label .text-wrapper-2 {
  color: var(--text-primary-basic);
  font-family: var(--paragraph-IBM-plex-sans-medium-font-family);
  font-size: var(--paragraph-IBM-plex-sans-medium-font-size);
  font-style: var(--paragraph-IBM-plex-sans-medium-font-style);
  font-weight: var(--paragraph-IBM-plex-sans-medium-font-weight);
  letter-spacing: var(--paragraph-IBM-plex-sans-medium-letter-spacing);
  line-height: var(--paragraph-IBM-plex-sans-medium-line-height);
  margin-top: -1px;
  position: relative;
  text-align: center;
  white-space: nowrap;
  width: fit-content;
}
</style>
