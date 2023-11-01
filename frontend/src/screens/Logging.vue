<template>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.1.1/css/all.min.css">
  <div class="logging">
    <div class="div-2">
      <div class="overlap-group">
        <p class="p">Welcome back. Please log in to your account</p>

        <text_input_component class="text-input-instance" state="default" text="Your email" text1="Type your email" />
        <text_input_component
          class="design-component-instance-node"
          state="password"
          text="Your password"
          text1="Type your password"
        />
        <div class="button-large-wrapper">
          <div class="button-large">
            <button_component class="button-large-filled" labelLabelClassName="button-instance" labelText="Log in" />
          </div>
        </div>
        <div class="div-wrapper">
          <div class="button-large-filled-wrapper">
            <button_component class="button-large-filled-2" labelLabelClassName="button-instance" labelText="Sign up" />
          </div>
        </div>
        <div class="text-wrapper-3">Forgot password?</div>
        <img class="image" alt="Image" src="../assets/img/image-3.png" />
      </div>
      <img class="img" alt="Image" src="../assets/img/image-2.png" />
      <div class="overlap">
        <div class="navigation">
          <nav_item_component class="nav-item-instance" />
          <nav_item_component class="nav-item-2" text="Accommodations" />
          <nav_item_component class="nav-item-3" text="Events" />
          <nav_item_component class="nav-item-4" text="Log in" />
          <nav_item_component class="nav-item-5" text="Activities" />
        </div>
        <img class="vector" alt="Vector" src="../assets/img/vector-1.svg" />
      </div>
    </div>
  </div>
</template>

<script>
import button_component from "../components/Button.vue";
import nav_item_component from "../components/NavItem.vue";
import text_input_component from "../components/TextInput.vue";
import '../assets/css/logging.css'

export default {
  name: "logging_page",
  components: {
    button_component,
    nav_item_component,
    text_input_component,
  },
};
</script>