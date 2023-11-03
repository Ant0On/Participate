<template>
  <div :class="['text-input', state, className]">
    <div class="div">
      <template v-if="state === 'active' || state === 'default' || state === 'error' || state === 'filled' || state === 'password' || state === 'read' || state === 'success'">
        {{ text }}
      </template>

      <template v-if="state === 'required'">
        <p class="span-wrapper">
          <input class="text-block-2" placeholder="Label">
        </p>
        <p class="span-wrapper">
          <span class="span">*</span>
        </p>
      </template>
    </div>
    <template v-if="['error', 'success'].includes(state)">
      <div class="text-block">
        <div class="text">
          <template v-if="state === 'error'">{{ text1 }}</template>

          <template v-if="state === 'success'">
            <p class="text-wrapper-4">A complex form might have multiple input fields</p>
          </template>
        </div>
      </div>
      <div class="error-message">
        <template v-if="state === 'error'">Error message</template>

        <template v-if="state === 'success'">Success message</template>
      </div>
    </template>

    <template v-if="['active', 'default', 'filled', 'password', 'read', 'required'].includes(state)">
      <div class="text-block-2">
        <template v-if="['active', 'default', 'filled', 'read', 'required'].includes(state)">
          <div class="text-2">
            <template v-if="['default', 'required'].includes(state)">
              <input class="text-block-2" :placeholder="text1"/>
            </template>

            <template v-if="['filled', 'read'].includes(state)">
              <p class="text-wrapper-4">
                A complex form might have multiple input fields, stacked vertically. Space the input fields evenly,
                clearly associating
              </p>
            </template>

            <template v-if="state === 'active'">A complex form might...|</template>
          </div>
        </template>

        <template v-if="['password'].includes(state)">
          <div class="text-3">
            <input class="text-block-2" :type="passwordFieldType" id="password" v-model="password" placeholder="**********">
            <i class="fas" :class="[passwordFieldIcon]" @click="hidePassword = !hidePassword"></i>
          </div>
        </template>
      </div>
    </template>
  </div>
</template>

<script>
import {computed, ref} from "vue";


export default {
  name: "text_input_component",
  props: {
    state: {
      type: String,
      validator: function (value) {
        return (
          [
            "active",
            "default",
            "success",
            "filled",
            "read",
            "date",
            "password",
            "error",
            "required",
          ].indexOf(value) !== -1
        );
      },
    },
    className: {
      type: String,
      default: "",
    },
    text: {
      type: String,
      default: "Label",
    },
    text1: {
      type: String,
      default: "Text",
    }
  },
  setup(props, context) {
    const hidePassword = ref(true);
    const password = ref("");

    const passwordFieldIcon = computed(() => hidePassword.value ? "fa-eye" : "fa-eye-slash");
    const passwordFieldType = computed(() => hidePassword.value ? "password" : "text");
    return {hidePassword, password, passwordFieldType, passwordFieldIcon}
  },
};
</script>

<style>
.text-input {
  align-items: flex-start;
  flex-direction: column;
  gap: 2px;
  position: relative;
}

.text-input .div {
  align-self: stretch;
  font-family: var(--paragraph-IBM-plex-sans-regular-font-family);
  font-size: var(--paragraph-IBM-plex-sans-regular-font-size);
  font-style: var(--paragraph-IBM-plex-sans-regular-font-style);
  font-weight: var(--paragraph-IBM-plex-sans-regular-font-weight);
  letter-spacing: var(--paragraph-IBM-plex-sans-regular-letter-spacing);
  line-height: var(--paragraph-IBM-plex-sans-regular-line-height);
  margin-top: -1px;
  position: relative;
}

.text-input .text-wrapper {
  color: #6f7482;
  font-family: var(--paragraph-IBM-plex-sans-regular-font-family);
  font-size: var(--paragraph-IBM-plex-sans-regular-font-size);
  font-style: var(--paragraph-IBM-plex-sans-regular-font-style);
  font-weight: var(--paragraph-IBM-plex-sans-regular-font-weight);
  letter-spacing: var(--paragraph-IBM-plex-sans-regular-letter-spacing);
  line-height: var(--paragraph-IBM-plex-sans-regular-line-height);
}

.text-input .span {
  color: #ed0131;
  font-family: var(--paragraph-IBM-plex-sans-regular-font-family);
  font-size: var(--paragraph-IBM-plex-sans-regular-font-size);
  font-style: var(--paragraph-IBM-plex-sans-regular-font-style);
  font-weight: var(--paragraph-IBM-plex-sans-regular-font-weight);
  letter-spacing: var(--paragraph-IBM-plex-sans-regular-letter-spacing);
  line-height: var(--paragraph-IBM-plex-sans-regular-line-height);
}

.text-input .text-block {
  align-items: flex-start;
  align-self: stretch;
  background-color: var(--surfacelight);
  border: 1px solid;
  border-radius: 4px;
  display: flex;
  flex: 0 0 auto;
  flex-direction: column;
  gap: 10px;
  padding: 12px;
  position: relative;
  width: 100%;
}

.text-input .text {
  align-self: stretch;
  color: var(--text-primary-basic);
  font-family: var(--paragraph-IBM-plex-sans-regular-font-family);
  font-size: var(--paragraph-IBM-plex-sans-regular-font-size);
  font-style: var(--paragraph-IBM-plex-sans-regular-font-style);
  font-weight: var(--paragraph-IBM-plex-sans-regular-font-weight);
  letter-spacing: var(--paragraph-IBM-plex-sans-regular-letter-spacing);
  line-height: var(--paragraph-IBM-plex-sans-regular-line-height);
  margin-top: -1px;
  position: relative;
}

.text-input .error-message {
  align-self: stretch;
  font-family: var(--smallest-descriptor-font-family);
  font-size: var(--smallest-descriptor-font-size);
  font-style: var(--smallest-descriptor-font-style);
  font-weight: var(--smallest-descriptor-font-weight);
  letter-spacing: var(--smallest-descriptor-letter-spacing);
  line-height: var(--smallest-descriptor-line-height);
  position: relative;
}

.text-input .text-block-2 {
  align-self: stretch;
  border-radius: 4px;
  position: relative;
  width: 100%;
}

.text-input .text-2 {
  align-self: stretch;
  color: var(--text-secondary-grey3);
  font-family: var(--paragraph-IBM-plex-sans-regular-font-family);
  font-size: var(--paragraph-IBM-plex-sans-regular-font-size);
  font-style: var(--paragraph-IBM-plex-sans-regular-font-style);
  font-weight: var(--paragraph-IBM-plex-sans-regular-font-weight);
  letter-spacing: var(--paragraph-IBM-plex-sans-regular-letter-spacing);
  line-height: var(--paragraph-IBM-plex-sans-regular-line-height);
  margin-top: -1px;
  position: relative;
}

.text-input .text-3 {
  align-self: stretch;
  color: var(--text-secondary-grey3);
  font-family: var(--paragraph-IBM-plex-sans-regular-font-family);
  font-size: var(--paragraph-IBM-plex-sans-regular-font-size);
  font-style: var(--paragraph-IBM-plex-sans-regular-font-style);
  font-weight: var(--paragraph-IBM-plex-sans-regular-font-weight);
  letter-spacing: var(--paragraph-IBM-plex-sans-regular-letter-spacing);
  line-height: var(--paragraph-IBM-plex-sans-regular-line-height);
  position: relative;
  margin-top: -1px;
  width: 400px;
}

.fa-eye {
  height: 24px;
  left: 424px;
  position: absolute;
  top: 12px;
  width: 24px;
}
.fa-eye-slash{
  height: 24px;
  left: 424px;
  position: absolute;
  top: 12px;
  width: 24px;
}

.text-input.active {
  display: flex;
  width: 460px;
}

.text-input.default {
  display: flex;
  width: 460px;
}

.text-input.read {
  display: inline-flex;
}

.text-input.date {
  display: inline-flex;
}

.text-input.filled {
  display: flex;
  width: 460px;
}

.text-input.success {
  display: flex;
  width: 460px;
}

.text-input.required {
  display: inline-flex;
}

.text-input.password {
  display: inline-flex;
}

.text-input.error {
  display: inline-flex;
}

.text-input.active .div {
  color: var(--text-secondary-grey2);
}

.text-input.default .div {
  color: var(--text-secondary-grey2);
}

.text-input.read .div {
  color: var(--text-secondary-grey2);
}

.text-input.date .div {
  color: var(--text-secondary-grey2);
}

.text-input.filled .div {
  color: var(--text-secondary-grey2);
}

.text-input.success .div {
  color: var(--text-secondary-grey2);
}

.text-input.required .div {
  color: transparent;
}

.text-input.password .div {
  color: var(--text-secondary-grey2);
}

.text-input.error .div {
  color: var(--text-secondary-grey2);
}

.text-input.active .text-block {
  border-color: var(--systemred);
}

.text-input.default .text-block {
  border-color: var(--systemred);
}

.text-input.read .text-block {
  border-color: var(--systemred);
}

.text-input.date .text-block {
  border-color: var(--systemred);
}

.text-input.filled .text-block {
  border-color: var(--systemred);
}

.text-input.success .text-block {
  border-color: var(--text-status-success);
}

.text-input.required .text-block {
  border-color: var(--systemred);
}

.text-input.password .text-block {
  border-color: var(--systemred);
}

.text-input.error .text-block {
  border-color: var(--systemred);
}

.text-input.active .error-message {
  color: var(--systemred);
}

.text-input.default .error-message {
  color: var(--systemred);
}

.text-input.read .error-message {
  color: var(--systemred);
}

.text-input.date .error-message {
  color: var(--systemred);
}

.text-input.filled .error-message {
  color: var(--systemred);
}

.text-input.success .error-message {
  color: var(--text-status-success);
}

.text-input.required .error-message {
  color: var(--systemred);
}

.text-input.password .error-message {
  color: var(--systemred);
}

.text-input.error .error-message {
  color: var(--systemred);
}

.text-input.active .text-block-2 {
  align-items: flex-start;
  background-color: var(--surfacelight);
  border: 1px solid;
  border-color: #0048d9;
  display: flex;
  flex: 0 0 auto;
  flex-direction: column;
  gap: 10px;
  padding: 12px;
}

.text-input.default .text-block-2 {
  background-color: var(--surfacelight);
  height: 48px;
}

.text-input.read .text-block-2 {
  align-items: flex-start;
  display: flex;
  flex: 0 0 auto;
  flex-direction: column;
  gap: 10px;
}

.text-input.date .text-block-2 {
  background-color: var(--surfacelight);
  height: 48px;
}

.text-input.filled .text-block-2 {
  align-items: flex-start;
  background-color: var(--surfacelight);
  display: flex;
  flex: 0 0 auto;
  flex-direction: column;
  gap: 10px;
  padding: 12px;
}

.text-input.required .text-block-2 {
  align-items: flex-start;
  background-color: var(--surfacelight);
  display: flex;
  flex: 0 0 auto;
  flex-direction: column;
  gap: 10px;
  padding: 12px;
}

.text-input.password .text-block-2 {
  background-color: var(--surfacelight);
  height: 48px;
}

.text-input.active .text-2 {
  color: var(--text-primary-basic);
}

.text-input.default .text-2 {
  color: var(--text-secondary-grey3);
}

.text-input.read .text-2 {
  color: var(--text-primary-basic);
}

.text-input.date .text-2 {
  color: var(--text-primary-basic);
}

.text-input.filled .text-2 {
  color: var(--text-primary-basic);
}

.text-input.success .text-2 {
  color: var(--text-primary-basic);
}

.text-input.required .text-2 {
  color: var(--text-secondary-grey3);
}

.text-input.password .text-2 {
  color: var(--text-primary-basic);
}

.text-input.error .text-2 {
  color: var(--text-primary-basic);
}
</style>
