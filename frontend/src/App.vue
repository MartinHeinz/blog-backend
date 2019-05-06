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
        <PostHeader title="This is title" author="Martin" published="05-05-2019 17:43"/>
        <Post text="Body of post"/>
        <BaseFooter :items="[
            {value: 'Home', url: '/home'},
            {value: 'Contact', url: '/contact'},
            {value: 'About', url: '/about'}]">
        </BaseFooter>
    </div>
</template>

<script>
import Navigation from './components/Navigation.vue';
import PostHeader from './components/PostHeader.vue';
import Post from './components/Post.vue';
import BaseFooter from './components/BaseFooter.vue';

export default {
    name: 'app',
    components: {
        Navigation,
        PostHeader,
        Post,
        BaseFooter,
    },
};
</script>
