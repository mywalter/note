<template>
  <div id="main">
    <menus></menus>
    <router-view></router-view>
    <modal></modal>
  </div>
</template>
<script>
  import menus from './menus';
  import modal from './modal';
  export default {
    components: {
      menus,
      modal
    }
  };
</script>

<style>
  * {
    padding: 0;
    margin: 0;
    font-family: "Open Sans", -apple-system, system-ui, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", "Microsoft Yahei", Arial, sans-serif;
  }

  a {
    text-decoration: none;
  }

  svg {
    width: 20px;
    height: 20px;
    vertical-align: -0.15em;
    fill: currentColor;
    overflow: hidden;
    color: #505f79;
  }

  input {
    height: auto;
    font-size: 20px;
    font-weight: 400;
    line-height: 1.3;
    width: 100%;
    transition: background-color 0.2s ease-in-out, border-color 0.2s ease-in-out;
    padding: 5px 7px;
    border-color: transparent;
    border-style: solid;
    border-width: 1px;
    border-radius: 4px;
    outline: none;
    background-color: transparent;
    -webkit-transition: background-color 0.2s ease-in-out, border-color 0.2s ease-in-out;
  }

  input:focus {
    background-color: #ffffff;
  }
  .color-66CCCC {
    color: #66CCCC
  }

  .color-FF99CC {
    color: #FF99CC
  }

  .color-FF9999 {
    color: #FF9999
  }

  .color-FFCC99 {
    color: #FFCC99
  }

  .color-FF6666 {
    color: #FF6666
  }

  .color-99CC66 {
    color: #99CC66
  }

  .color-666699 {
    color: #666699
  }

  .color-FF9999 {
    color: #FF9999
  }

  .color-99CC33 {
    color: #99CC33
  }

  .color-FF9900 {
    color: #FF9900
  }

  .color-FFCC00 {
    color: #FFCC00
  }

  .color-FF0033 {
    color: #FF0033
  }

  .color-FF9966 {
    color: #FF9966
  }

  .color-FF9900 {
    color: #FF9900
  }

  .color-CC3399 {
    color: #CC3399
  }

  .color-99CC33 {
    color: #99CC33
  }

  .color-FF6600 {
    color: #FF6600
  }

  .color-993366 {
    color: #993366
  }

  .color-CCCC33 {
    color: #CCCC33
  }

  .color-666633 {
    color: #666633
  }

  .color-66CCCC {
    color: #66CCCC
  }

  .color-666699 {
    color: #666699
  }

  .color-CC9999 {
    color: #CC9999
  }

  .color-666699 {
    color: #666699
  }

  .color-FF9900 {
    color: #FF9900
  }

  .color-0099CC {
    color: #0099CC
  }

  .color-CCCC99 {
    color: #CCCC99
  }

  .color-CC3399 {
    color: #CC3399
  }

  .color-99CC00 {
    color: #99CC00
  }

  .color-FF6666 {
    color: #FF6666
  }

  .color-3399CC {
    color: #3399CC
  }

  .color-CC6600 {
    color: #CC6600
  }

  .color-999999 {
    color: #999999
  }

  .color-CCCC33 {
    color: #CCCC33
  }

  .color-FF9933 {
    color: #FF9933
  }

  .color-009933 {
    color: #009933
  }

  .color-0099CC {
    color: #0099CC
  }

  .color-CCCCCC {
    color: #CCCCCC
  }

  .color-FF6666 {
    color: #FF6666
  }

  .color-FF6600 {
    color: #FF6600
  }

  .color-009966 {
    color: #009966
  }

  .color-CC6633 {
    color: #CC6633
  }

  .color-FFCC99 {
    color: #FFCC99
  }

  .color-CC6600 {
    color: #CC6600
  }

  .color-CC0066 {
    color: #CC0066
  }

  .color-009999 {
    color: #009999
  }

  .color-FFCC33 {
    color: #FFCC33
  }

  #main {
    display: flex;
    position: relative;
    flex-grow: 1;
  }
</style>
