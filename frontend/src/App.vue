<template>
    <div id="app">
        <Navigation
                :items="[
            {value: 'Home', url: '/home'},
            {value: 'Contact', url: '/contact'},
            {value: 'About', url: '/about'}]"
                :actions="{
            url_previous: 'url_previous_test',
            url_next: 'url_previous_test',
                }">
        </Navigation>

        <Post info="Welcome to Your Vue.js App"/>
        <BaseFooter :items="[
            {value: 'Home', url: '/home'},
            {value: 'Contact', url: '/contact'},
            {value: 'About', url: '/about'}]">
        </BaseFooter>
    </div>
</template>

<script>
import Post from './components/Post.vue';
import BaseFooter from './components/BaseFooter.vue';
import Navigation from './components/Navigation.vue';

export default {
    name: 'app',
    components: {
        Navigation,
        Post,
        BaseFooter,
    },
};
</script>
