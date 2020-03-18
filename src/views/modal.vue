<template>
  <div id="modal">
    <div class="mask"></div>
    <div class="modal">
      <div class="info">
        <div class="infoHeader">
          <div class="title">
            magento
          </div>
          <div class="iconRight">
            <i class="iconfont icon-arrow-right"></i>
          </div>
        </div>
        <div class="infoBody">
          <form action="">
            <div class="infoGroup">
              <div class="label">
                <span>标题</span>
              </div>
              <div class="form-control">
                <input type="text" class="inputInner">
              </div>
            </div>
            <div class="infoGroup">
              <div class="label">
                <span>UG</span>
              </div>
              <div class="form-control">
                <input type="text" class="inputInner">
              </div>
            </div>
            <div class="infoButtons">
              <div class="buttons">
                <div class="actionsSecondary">
                  <button>
                                <span>
                                    取消
                                </span>
                  </button>
                </div>
                <div class="actionsPrimary">
                  <button>
                                <span>
                                    保存
                                </span>
                  </button>
                </div>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

</script>
<style>
  #modal .mask {
    top: 0px;
    left: 0px;
    right: 0px;
    bottom: 0px;
    display: flex;
    z-index: 30;
    position: fixed;
    align-items: center;
    animation-name: none;
    pointer-events: none;
    justify-content: center;
    -webkit-box-pack: center;
    background-color: rgba(24, 48, 85, 0);
    -webkit-box-align: center;
    animation-duration: 0.2s;
    animation-fill-mode: forwards;
    animation-timing-function: ease-out;
    margin: 0px;
    padding: 0px;
    transition: background-color 250ms ease-out 0s;
  }

  #modal .search {
    top: 0px;
    left: 100%;
    width: 385px;
    bottom: 0px;
    display: flex;
    position: absolute;
    align-self: stretch;
    align-items: stretch;
    flex-direction: column;
    pointer-events: all;
    -webkit-box-align: stretch;
    -webkit-box-orient: vertical;
    -webkit-box-direction: normal;
    margin: 0px;
    padding: 0px;
    z-index: 9999;
    background: rgb(255, 255, 255);
    transition: transform 250ms ease 0s;
  }

  #modal .searchHeader {
    width: 100%;
    height: 80px;
    display: flex;
    max-width: 385px;
    align-items: center;
    flex-direction: row;
    -webkit-box-align: center;
    -webkit-box-orient: horizontal;
    -webkit-box-direction: normal;
    margin: 0px;
    padding: 0px;
    border-bottom: 1px solid rgb(230, 236, 241);
  }

  #modal .searchHeader .iconSearch {
    color: rgb(59, 69, 78);
    cursor: pointer;
    height: 48px;
    display: inline-flex;
    align-items: center;
    line-height: 1em;
    white-space: nowrap;
    justify-content: center;
    -webkit-box-pack: center;
    background-color: transparent;
    -webkit-box-align: center;
    border-width: 1px;
    border-style: solid;
    border-image: initial;
    outline: none;
    padding-left: 24px;
    border-color: transparent;
    border-radius: 3px;
    text-decoration: none;
    transition: all 250ms ease-out 0s;
  }

  #modal .searchHeader .iconRight {
    color: rgb(59, 69, 78);
    width: 48px;
    cursor: pointer;
    height: 48px;
    display: inline-flex;
    align-items: center;
    line-height: 1em;
    white-space: nowrap;
    justify-content: center;
    -webkit-box-pack: center;
    background-color: transparent;
    -webkit-box-align: center;
    border-width: 1px;
    border-style: solid;
    border-image: initial;
    outline: none;
    padding: 0px;
    border-color: transparent;
    border-radius: 3px;
    text-decoration: none;
    transition: all 250ms ease-out 0s;
  }

  #modal .searchHeader .iconfont {
    font-size: 14px;
  }

  #modal .searchHeader .title {
    color: rgb(36, 42, 49);
    display: block;
    flex: 1 1 0%;
    margin: 0px;
  }

  #modal .searchHeader .title input {
    color: rgb(36, 42, 49);
    width: 100%;
    cursor: inherit;
    height: 38px;
    resize: none;
    box-sizing: border-box;
    text-align: left;
    font: inherit;
    border-width: initial;
    border-style: none;
    border-color: initial;
    border-image: initial;
    margin: 0px;
    outline: none;
    padding: 0px 8px;
    background: transparent;
    border-radius: 3px;
  }

  #modal .searchBody {
    width: 100%;
    display: block;
    max-width: 385px;
    overflow-x: hidden;
    overflow-y: overlay;
    flex: 1 1 0%;
    margin: 0px;
    padding: 0px;
  }

  #modal .searchBody .searchContent {
    width: 100%;
    height: 100%;
    display: flex;
    flex: 1 1 0%;
    margin: 0px;
    padding: 0px;
  }

  #modal .searchBody .searchContent .searchContentNull {
    display: flex;
    align-items: center;
    flex-direction: column;
    justify-content: center;
    -webkit-box-pack: center;
    -webkit-box-align: center;
    -webkit-box-orient: vertical;
    -webkit-box-direction: normal;
    flex: 1 1 0%;
    margin: 0px;
    padding: 16px 16px 56px;
  }

  #modal .searchBody .searchContent .searchContentNull .searchImage {
    display: block;
    margin: 0px 0px 24px;
    padding: 0px;
  }

  #modal .searchBody .searchContent .searchContentNull .searchMessage {
    color: rgb(157, 170, 182);
    display: block;
    text-align: center;
    margin: 0px;
    padding: 0px;
  }

  #modal .searchBody .searchContent .searchContentNull .searchMessage span {
    font-size: 16px;
    font-family: Roboto, sans-serif;
    font-weight: 500;
    line-height: 1.5;
  }

  #modal .searchBody .searchBodyList {
    display: block;
    margin: 0px;
    padding: 0px;
  }

  #modal .searchBody .searchBodyList .searchItemUl {
    cursor: pointer;
    display: block;
    margin: 0px;
    padding: 0px;
  }

  #modal .searchBody .searchBodyList .searchItem {
    display: block;
    margin: 0px;
    padding: 16px 24px;
    border-bottom: 1px solid rgb(230, 236, 241);
  }

  #modal .searchBody .searchBodyList .searchItem .searchItemTitle {
    color: rgb(59, 69, 78);
    display: block;
    font-size: 14px;
    font-weight: 500;
    line-height: 1.5;
    white-space: nowrap;
    text-overflow: ellipsis;
    margin: 0px;
    padding: 0px 0px 8px;
    overflow: hidden;
  }

  #modal .searchBody .searchBodyList .searchItem .searchItemContent {
    color: rgb(116, 129, 141);
    font-size: 12px;
    font-family: Roboto, sans-serif;
    font-weight: 400;
    line-height: 1.625;
  }

  #modal .info {
    top: 0px;
    left: 100%;
    width: 385px;
    bottom: 0px;
    display: flex;
    position: absolute;
    align-self: stretch;
    align-items: stretch;
    flex-direction: column;
    pointer-events: all;
    -webkit-box-align: stretch;
    -webkit-box-orient: vertical;
    -webkit-box-direction: normal;
    margin: 0px;
    padding: 0px;
    z-index: 9999;
    background: rgb(255, 255, 255);
    transition: transform 250ms ease 0s;
  }

  #modal .info .infoHeader {
    height: 80px;
    display: flex;
    align-items: center;
    -webkit-box-align: center;
    padding-left: 24px;
    padding-right: 24px;
    margin: 0px;
    border-bottom: 1px solid rgb(230, 236, 241);
  }

  #modal .info .infoHeader .title {
    color: rgb(116, 129, 141);
    font-size: 12px;
    font-family: Roboto, sans-serif;
    font-weight: 700;
    line-height: 1.2;
    letter-spacing: 1.2px;
    text-transform: uppercase;
    flex: 1 1 0%;
  }

  #modal .info .infoBody {
    display: flex;
    overflow-x: hidden;
    overflow-y: overlay;
    flex-direction: column;
    -webkit-box-orient: vertical;
    -webkit-box-direction: normal;
    flex: 1 1 0%;
    margin: 0px;
    padding: 24px;
  }

  #modal .info .infoBody form {
    display: flex;
    flex-direction: column;
    -webkit-box-orient: vertical;
    -webkit-box-direction: normal;
    flex: 1 1 auto;
  }

  #modal .info .infoBody form .infoGroup {
    display: block;
    text-align: left;
    margin: 0px 0px 24px;
    padding: 0px;
  }

  #modal .info .infoBody form .infoGroup .label {
    display: flex;
    align-items: center;
    justify-content: space-between;
    -webkit-box-pack: justify;
    -webkit-box-align: center;
    margin: 0px 0px 8px;
    padding: 0px;
  }

  #modal .info .infoBody form .infoGroup .label span {
    color: rgb(157, 170, 182);
    font-size: 12px;
    font-family: Roboto, sans-serif;
    font-weight: 700;
    line-height: 1.2;
    letter-spacing: 1.2px;
    text-transform: uppercase;
  }

  #modal .info .infoBody form .infoGroup .form-control {
    width: 100%;
    display: flex;
    font-size: 16px;
    box-sizing: border-box;
    font-weight: 400;
    line-height: 1.625;
    border-width: 1px;
    border-style: solid;
    border-image: initial;
    margin: 0px;
    outline: none;
    padding: 0px;
    background: rgb(255, 255, 255);
    border-color: rgb(211, 220, 228);
    border-radius: 3px;
  }

  #modal .info .infoBody form .infoGroup .form-control .inputInner {
    color: rgb(36, 42, 49);
    width: 100%;
    cursor: inherit;
    height: 38px;
    resize: none;
    box-sizing: border-box;
    text-align: left;
    font: inherit;
    border-width: initial;
    border-style: none;
    border-color: initial;
    border-image: initial;
    margin: 0px;
    outline: none;
    padding: 0px 8px;
    background: transparent;
    border-radius: 3px;
  }

  #modal .info .infoBody form .infoButtons {
    display: block;
    margin: 0px;
    padding: 24px 0px;
    border-top: 1px solid rgb(230, 236, 241);
  }

  #modal .info .infoBody form .infoButtons .buttons {
    display: flex;
    margin: 0px;
    padding: 0px;
  }

  #modal .info .infoBody form .infoButtons .buttons .actionsSecondary {
    display: block;
    margin: 0px;
    padding: 0px;
  }

  #modal .info .infoBody form .infoButtons .buttons .actionsPrimary {
    display: flex;
    justify-content: flex-end;
    -webkit-box-pack: end;
    flex: 1 1 0%;
    margin: 0px;
    padding: 0px;
  }

  #modal .info .infoBody form .infoButtons .buttons .actionsSecondary button {
    color: rgb(116, 129, 141);
    cursor: pointer;
    height: 40px;
    display: inline-flex;
    align-items: center;
    line-height: 1em;
    white-space: nowrap;
    justify-content: center;
    -webkit-box-pack: center;
    background-color: rgb(230, 236, 241);
    -webkit-box-align: center;
    border-width: 1px;
    border-style: solid;
    border-image: initial;
    outline: none;
    padding: 0px 24px;
    border-color: transparent;
    border-radius: 3px;
    text-decoration: none;
    transition: all 250ms ease-out 0s;
  }

  #modal .info .infoBody form .infoButtons .buttons .actionsPrimary button {
    color: rgb(255, 255, 255);
    cursor: pointer;
    height: 40px;
    display: inline-flex;
    align-items: center;
    line-height: 1em;
    white-space: nowrap;
    justify-content: center;
    -webkit-box-pack: center;
    background-color: rgb(56, 132, 255);
    -webkit-box-align: center;
    border-width: 1px;
    border-style: solid;
    border-image: initial;
    outline: none;
    padding: 0px 24px;
    border-color: transparent;
    border-radius: 3px;
    text-decoration: none;
    transition: all 250ms ease-out 0s;
  }

  #modal .info .infoBody form .infoButtons .buttons button span {
    font-size: 16px;
    font-family: Roboto, sans-serif;
    font-weight: 500;
    line-height: 1em;
  }
</style>