<template>
    <div>
        <Menu tagIndex="1"></Menu>
        <div class="rightMain">
            <div style="padding:0 0 0 0;height:60px;margin-bottom:20px;">
                <div style="background-color: #ffffff;height:60px;padding:10px;"></div>
            </div>
            <div id="list">
                <div style="margin:10px;">主题</div>
            </div>
        </div>
    </div>
</template>

<script>
  import Menu from './components/menu';

  export default {
    data: function () {
      return {
      }
    },
    components:{
        Menu
    }
  }
</script>