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
        <div class="content">
            <PostHeader title="This is title"
                        author="Martin"
                        published="05-05-2019 17:43"
                        :tags="[
                            {name: 'Golang', url: 'go'},
                            {name: 'Vue.js', url: 'vue'}
                        ]">
            </PostHeader>
            <Post text="Body of post"/>
            <BaseFooter :items="[
                {value: 'Home', url: '/home'},
                {value: 'Contact', url: '/contact'},
                {value: 'About', url: '/about'}]">
            </BaseFooter>
        </div>
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

<style>
    body {
        max-width: 48rem;
        margin: 0;
        height: 100%;
        background-color: #1d1f21;
        color: #c9cacc;
        font-weight: 400;
        font-size: 14px;
        font-family: "Menlo", "Meslo LG", monospace;
        line-height: 1.725;
        text-rendering: geometricPrecision;
        flex: 1;
        -moz-osx-font-smoothing: grayscale;
        -webkit-font-smoothing: antialiased;
        padding-right: 2rem;
        padding-left: 2rem;
    }

    .content {
        position: relative;
        display: flex;
        flex-direction: column;
        min-height: 100%;
        overflow-wrap: break-word;
        padding-top: 4rem;
        padding-bottom: 4rem;
        box-sizing: border-box;
    }
</style>
