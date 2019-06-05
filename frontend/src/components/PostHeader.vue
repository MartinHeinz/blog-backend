<template>
    <header>

        <h1 class="posttitle" itemprop="name headline">{{title}}</h1>


        <div class="meta">

            <span class="author" itemprop="author" itemscope="" itemtype="http://schema.org/Person">
                <span itemprop="name">{{author}}</span>
            </span>

            <div class="postdate">
                <time :datetime="published | formatDate" itemprop="datePublished">{{published}}</time>
            </div>

            <div class="article-tag">
                <v-icon>fas fa-tag</v-icon>
                <BaseTag v-for="tag in tags"
                              v-bind:url="tag.url"
                              v-bind:name="tag.name"
                              v-bind:key="tag.key">
                </BaseTag>
            </div>

        </div>
    </header>
</template>

<script>
import BaseTag from './BaseTag.vue';

export default {
    name: 'PostHeader',
    components: { BaseTag },
    props: {
        tags: Array,
        title: String,
        author: String,
        published: String,
    },
};
</script>

<style>

    .article-tag .tag-link {
        background-image: linear-gradient(transparent, transparent 10px, #d480aa 10px, #d480aa);
        background-position: bottom;
        background-size: 100% 6px;
        background-repeat: repeat-x;
    }


    .tag-link a:hover {
        box-shadow: inset 0 0px 0 white, inset 0 -2px 0 #d480aa;
    }

    .tag-link a {
        color: #c9cacc;
        text-decoration: none;
        background-image: linear-gradient(transparent, transparent 5px, #c9cacc 5px, #c9cacc);
        background-position: bottom;
        background-size: 100% 6px;
        background-repeat: repeat-x;
    }

    article header .posttitle {
        margin-top: 0;
        margin-bottom: 0;
        text-transform: none;
        font-size: 1.5em;
        line-height: 1.25;
    }

    .posttitle {
        text-transform: none;
        font-size: 1.5em;
        line-height: 1.25;
    }

    header .meta {
        margin-top: 0;
        margin-bottom: 1rem;
    }
    header .meta * {
        color: #ccc;
        font-size: 0.85rem;
    }

    header .author {
        text-transform: uppercase;
        letter-spacing: 0.01em;
        font-weight: 700;
        margin: 0 5px;
    }

    header .postdate {
        display: inline;
    }

    @media (min-width: 480px) {
        .article-tag {
            display: inline;
        }
        .article-tag:before {
            content: " | ";
        }
    }

    span:first-child {
        margin-left:0;
    }
</style>
