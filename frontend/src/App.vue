<template>
    <div id="app">
        <BaseMenu
            :items="[
            {value: 'Home', url: '/home'},
            {value: 'Contact', url: '/contact'},
            {value: 'About', url: '/about'}]">
        </BaseMenu>
        <Post info="Welcome to Your Vue.js App"/>
        <Footer/>
    </div>
</template>

<script>
import Post from './components/Post.vue';
import Footer from './components/Footer.vue';
import BaseMenu from './components/BaseMenu.vue';

export default {
    name: 'app',
    components: {
        BaseMenu,
        Post,
        Footer,
    },
};
</script>
