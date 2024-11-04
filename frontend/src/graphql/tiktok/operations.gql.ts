import { gql } from '@urql/vue';

const TIKTOK_FRAGMENT = gql`
  fragment TiktokFragment on VideoItem {
    type
    common {
      doc_id_str
    }
    item {
      id
      desc
      originalItem
      officalItem
      secret
      forFriend
      secret
      digged
      itemCommentStatus
      showNotPass
      itemMute
      privateItem
      duetEnabled
      stitchEnabled
      shareEnabled
      authorStats {
        followingCount
        followerCount
        heartCount
        videoCount
        diggCount
        heart
      }
      textExtra {
        awemeId
        start
        end
        hashtagName
        hashtagId
        type
      }
      stats {
        diggCount
        shareCount
        commentCount
        playCount
        collectCount
      }
      challenges {
        id
        title
        desc
        profileThumb
        profileMedium
        profileLarger
        coverThumb
        coverMedium
        coverLarger
      }
      music {
        id
        title
        playUrl
        coverThumb
        coverMedium
        coverLarge
        authorName
        original
        duration
        album
      }
      author {
        id
        uniqueId
        nickname
        avatarThumb
        avatarMedium
        avatarLarger
        signature
        verified
      }
      video {
        id
        height
        width
        duration
        ratio
        cover
        originCover
        dynamicCover
        playAddr
        playAddrPublic
        reflowCover
        bitrate
        encodedType
        format
        videoQuality
        encodeUserTag
        shareCover
      }
    }
  }
`

export const TIKTOK_SEARCH = gql`
	${TIKTOK_FRAGMENT}

	query searchTiktok($input: String!) {
		searchTiktok(keyword: $input) {
      ...TiktokFragment
		}
	}
`;

export const TIKTOK_GET_BY_ID = gql`
	${TIKTOK_FRAGMENT}

	query getTiktok($input: String!) {
		getTiktok(id: $input) {
      ...TiktokFragment
		}
	}
`;

export const TIKTOK_SUB_SEARCH = gql`
	${TIKTOK_FRAGMENT}

	subscription subSearchTiktok($input: String!) {
		subSearchTiktok(keyword: $input) {
      ...TiktokFragment
		}
	}
`;
