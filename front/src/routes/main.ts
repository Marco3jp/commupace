import Vue from 'vue';
import Router from 'vue-router';

import Community from '@view/community';
import CommunityFallback from '@components/community/communityFallBack.vue';
import CommunityFrame from '@view/frame/communityFrame.vue';
import CommunityFront from '@components/community/communityFront.vue';
import CommunityChat from '@components/community/chat/communityChat.vue';
import CommunityThreadFrame from '@view/frame/communityThreadFrame';
import CommunityThreadFront from '@components/community/thread/communityThreadList.vue';
import CommunityThread from '@components/community/thread/communityThread.vue';
import Explore from '@view/explore';
import MyPage from '@view/mypage';
import signUp from "@view/signUp.vue";

Vue.use(Router);

/**
 * path: "urlパス"
 * component: "使用しているコンポーネント"
 */

export default new Router({
    mode: 'history',
    routes: [
        {
            path: '/signUp',
            name: 'signUp',
            component: signUp,
        },
        {
            path: '/myPage',
            name: 'myPage',
            component: MyPage
        },
        {
            path: '/community',
            component: Community,
            children: [
                {
                    path: '',
                    component: CommunityFallback,
                    name: 'community',
                },
                {
                    path: ':communityId',
                    component: CommunityFrame,
                    children: [
                        {
                            path: '',
                            name: 'communityFront',
                            component: CommunityFront,
                        },
                        {
                            path: 'chat',
                            name: 'communityChat',
                            component: CommunityChat,
                        },
                        {
                            path: 'thread',
                            component: CommunityThreadFrame,
                            children: [
                                {
                                    path: '',
                                    name: 'communityThreadFront',
                                    component: CommunityThreadFront,
                                },
                                {
                                    path: ':threadId',
                                    name: 'communityThread',
                                    component: CommunityThread,
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            path: '/explore',
            name: 'explore',
            component: Explore
        },
    ]
});