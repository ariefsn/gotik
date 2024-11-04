import { z } from 'zod'
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Map: { input: any; output: any; }
  Time: { input: any; output: any; }
  Upload: { input: any; output: any; }
};

export type Audit = {
  __typename?: 'Audit';
  createdAt: Scalars['Time']['output'];
  createdBy?: Maybe<Scalars['String']['output']>;
  publishedAt?: Maybe<Scalars['Time']['output']>;
  publishedBy?: Maybe<Scalars['String']['output']>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['String']['output']>;
};

export type Author = {
  __typename?: 'Author';
  avatarLarger?: Maybe<Scalars['String']['output']>;
  avatarMedium?: Maybe<Scalars['String']['output']>;
  avatarThumb?: Maybe<Scalars['String']['output']>;
  commentSetting?: Maybe<Scalars['Int']['output']>;
  downloadSetting?: Maybe<Scalars['Int']['output']>;
  duetSetting?: Maybe<Scalars['Int']['output']>;
  ftc?: Maybe<Scalars['Boolean']['output']>;
  id?: Maybe<Scalars['String']['output']>;
  nickname?: Maybe<Scalars['String']['output']>;
  openFavorite?: Maybe<Scalars['Boolean']['output']>;
  privateAccount?: Maybe<Scalars['Boolean']['output']>;
  relation?: Maybe<Scalars['Int']['output']>;
  secUid?: Maybe<Scalars['String']['output']>;
  secret?: Maybe<Scalars['Boolean']['output']>;
  signature?: Maybe<Scalars['String']['output']>;
  stitchSetting?: Maybe<Scalars['Int']['output']>;
  uniqueId?: Maybe<Scalars['String']['output']>;
  verified?: Maybe<Scalars['Boolean']['output']>;
};

export type AuthorStats = {
  __typename?: 'AuthorStats';
  diggCount?: Maybe<Scalars['Int']['output']>;
  followerCount?: Maybe<Scalars['Int']['output']>;
  followingCount?: Maybe<Scalars['Int']['output']>;
  heart?: Maybe<Scalars['Int']['output']>;
  heartCount?: Maybe<Scalars['Int']['output']>;
  videoCount?: Maybe<Scalars['Int']['output']>;
};

export type Challenge = {
  __typename?: 'Challenge';
  coverLarger?: Maybe<Scalars['String']['output']>;
  coverMedium?: Maybe<Scalars['String']['output']>;
  coverThumb?: Maybe<Scalars['String']['output']>;
  desc?: Maybe<Scalars['String']['output']>;
  id?: Maybe<Scalars['String']['output']>;
  isCommerce?: Maybe<Scalars['Boolean']['output']>;
  profileLarger?: Maybe<Scalars['String']['output']>;
  profileMedium?: Maybe<Scalars['String']['output']>;
  profileThumb?: Maybe<Scalars['String']['output']>;
  title?: Maybe<Scalars['String']['output']>;
};

export type Common = {
  __typename?: 'Common';
  doc_id_str?: Maybe<Scalars['String']['output']>;
};

export type DuetInfo = {
  __typename?: 'DuetInfo';
  duetFromId?: Maybe<Scalars['String']['output']>;
};

export type Extra = {
  __typename?: 'Extra';
  api_debug_info?: Maybe<Scalars['String']['output']>;
  fatal_item_ids?: Maybe<Array<Maybe<Scalars['String']['output']>>>;
  logid?: Maybe<Scalars['String']['output']>;
  now?: Maybe<Scalars['Int']['output']>;
  search_request_id?: Maybe<Scalars['String']['output']>;
};

export type GlobalDoodleConfig = {
  __typename?: 'GlobalDoodleConfig';
  feedback_survey?: Maybe<Scalars['String']['output']>;
};

export type Item = {
  __typename?: 'Item';
  author?: Maybe<Author>;
  authorStats?: Maybe<AuthorStats>;
  challenges?: Maybe<Array<Maybe<Challenge>>>;
  collected?: Maybe<Scalars['Boolean']['output']>;
  createTime?: Maybe<Scalars['Int']['output']>;
  desc?: Maybe<Scalars['String']['output']>;
  digged?: Maybe<Scalars['Boolean']['output']>;
  duetEnabled?: Maybe<Scalars['Boolean']['output']>;
  duetInfo?: Maybe<DuetInfo>;
  forFriend?: Maybe<Scalars['Boolean']['output']>;
  id?: Maybe<Scalars['String']['output']>;
  isAd?: Maybe<Scalars['Boolean']['output']>;
  itemCommentStatus?: Maybe<Scalars['Int']['output']>;
  itemMute?: Maybe<Scalars['Boolean']['output']>;
  music?: Maybe<Music>;
  officalItem?: Maybe<Scalars['Boolean']['output']>;
  originalItem?: Maybe<Scalars['Boolean']['output']>;
  privateItem?: Maybe<Scalars['Boolean']['output']>;
  secret?: Maybe<Scalars['Boolean']['output']>;
  shareEnabled?: Maybe<Scalars['Boolean']['output']>;
  showNotPass?: Maybe<Scalars['Boolean']['output']>;
  stats?: Maybe<Stats>;
  stitchEnabled?: Maybe<Scalars['Boolean']['output']>;
  textExtra?: Maybe<Array<Maybe<TextExtra>>>;
  video?: Maybe<Video>;
  vl1?: Maybe<Scalars['Boolean']['output']>;
};

export type LogPb = {
  __typename?: 'LogPb';
  impr_id?: Maybe<Scalars['String']['output']>;
};

export type Music = {
  __typename?: 'Music';
  album?: Maybe<Scalars['String']['output']>;
  authorName?: Maybe<Scalars['String']['output']>;
  coverLarge?: Maybe<Scalars['String']['output']>;
  coverMedium?: Maybe<Scalars['String']['output']>;
  coverThumb?: Maybe<Scalars['String']['output']>;
  duration?: Maybe<Scalars['Int']['output']>;
  id?: Maybe<Scalars['String']['output']>;
  original?: Maybe<Scalars['Boolean']['output']>;
  playUrl?: Maybe<Scalars['String']['output']>;
  title?: Maybe<Scalars['String']['output']>;
};

export type Mutation = {
  __typename?: 'Mutation';
  createTodo: Todo;
  updateTodo: Todo;
};


export type MutationCreateTodoArgs = {
  input: NewTodoDto;
};


export type MutationUpdateTodoArgs = {
  input: UpdateTodoDto;
};

export type NewTodoDto = {
  text: Scalars['String']['input'];
  userId: Scalars['String']['input'];
};

export type Query = {
  __typename?: 'Query';
  getTiktok?: Maybe<VideoItem>;
  searchTiktok?: Maybe<Array<VideoItem>>;
  todos: Array<Todo>;
};


export type QueryGetTiktokArgs = {
  id: Scalars['String']['input'];
};


export type QuerySearchTiktokArgs = {
  keyword: Scalars['String']['input'];
};

export type Stats = {
  __typename?: 'Stats';
  collectCount?: Maybe<Scalars['Int']['output']>;
  commentCount?: Maybe<Scalars['Int']['output']>;
  diggCount?: Maybe<Scalars['Int']['output']>;
  playCount?: Maybe<Scalars['Int']['output']>;
  shareCount?: Maybe<Scalars['Int']['output']>;
};

export type Subscription = {
  __typename?: 'Subscription';
  onNewTodo?: Maybe<Todo>;
  onUpdateTodo?: Maybe<Todo>;
  subSearchTiktok?: Maybe<Array<VideoItem>>;
};


export type SubscriptionSubSearchTiktokArgs = {
  keyword: Scalars['String']['input'];
};

export type TextExtra = {
  __typename?: 'TextExtra';
  awemeId?: Maybe<Scalars['String']['output']>;
  end?: Maybe<Scalars['Int']['output']>;
  hashtagId?: Maybe<Scalars['String']['output']>;
  hashtagName?: Maybe<Scalars['String']['output']>;
  isCommerce?: Maybe<Scalars['Boolean']['output']>;
  secUid?: Maybe<Scalars['String']['output']>;
  start?: Maybe<Scalars['Int']['output']>;
  subType?: Maybe<Scalars['Int']['output']>;
  type?: Maybe<Scalars['Int']['output']>;
  userId?: Maybe<Scalars['String']['output']>;
  userUniqueId?: Maybe<Scalars['String']['output']>;
};

export type TiktokSearchResponse = {
  __typename?: 'TiktokSearchResponse';
  backtrace?: Maybe<Scalars['String']['output']>;
  cursor?: Maybe<Scalars['Int']['output']>;
  data?: Maybe<Array<Maybe<VideoItem>>>;
  extra?: Maybe<Extra>;
  global_doodle_config?: Maybe<GlobalDoodleConfig>;
  has_more?: Maybe<Scalars['Int']['output']>;
  log_pb?: Maybe<LogPb>;
  qc?: Maybe<Scalars['String']['output']>;
  status_code?: Maybe<Scalars['Int']['output']>;
};

export type Todo = {
  __typename?: 'Todo';
  done: Scalars['Boolean']['output'];
  id: Scalars['ID']['output'];
  text: Scalars['String']['output'];
  user: User;
};

export type UpdateTodoDto = {
  done: Scalars['Boolean']['input'];
  id: Scalars['ID']['input'];
  text: Scalars['String']['input'];
};

export type User = {
  __typename?: 'User';
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
};

export type Video = {
  __typename?: 'Video';
  bitrate?: Maybe<Scalars['Int']['output']>;
  cover?: Maybe<Scalars['String']['output']>;
  duration?: Maybe<Scalars['Int']['output']>;
  dynamicCover?: Maybe<Scalars['String']['output']>;
  encodeUserTag?: Maybe<Scalars['String']['output']>;
  encodedType?: Maybe<Scalars['String']['output']>;
  format?: Maybe<Scalars['String']['output']>;
  height?: Maybe<Scalars['Int']['output']>;
  id?: Maybe<Scalars['String']['output']>;
  originCover?: Maybe<Scalars['String']['output']>;
  playAddr?: Maybe<Scalars['String']['output']>;
  playAddrPublic?: Maybe<Scalars['String']['output']>;
  ratio?: Maybe<Scalars['String']['output']>;
  reflowCover?: Maybe<Scalars['String']['output']>;
  shareCover?: Maybe<Array<Maybe<Scalars['String']['output']>>>;
  videoQuality?: Maybe<Scalars['String']['output']>;
  width?: Maybe<Scalars['Int']['output']>;
};

export type VideoItem = {
  __typename?: 'VideoItem';
  common?: Maybe<Common>;
  item?: Maybe<Item>;
  type?: Maybe<Scalars['Int']['output']>;
};

export type TiktokFragmentFragment = { __typename?: 'VideoItem', type?: number | null, common?: { __typename?: 'Common', doc_id_str?: string | null } | null, item?: { __typename?: 'Item', id?: string | null, desc?: string | null, originalItem?: boolean | null, officalItem?: boolean | null, secret?: boolean | null, forFriend?: boolean | null, digged?: boolean | null, itemCommentStatus?: number | null, showNotPass?: boolean | null, itemMute?: boolean | null, privateItem?: boolean | null, duetEnabled?: boolean | null, stitchEnabled?: boolean | null, shareEnabled?: boolean | null, authorStats?: { __typename?: 'AuthorStats', followingCount?: number | null, followerCount?: number | null, heartCount?: number | null, videoCount?: number | null, diggCount?: number | null, heart?: number | null } | null, textExtra?: Array<{ __typename?: 'TextExtra', awemeId?: string | null, start?: number | null, end?: number | null, hashtagName?: string | null, hashtagId?: string | null, type?: number | null } | null> | null, stats?: { __typename?: 'Stats', diggCount?: number | null, shareCount?: number | null, commentCount?: number | null, playCount?: number | null, collectCount?: number | null } | null, challenges?: Array<{ __typename?: 'Challenge', id?: string | null, title?: string | null, desc?: string | null, profileThumb?: string | null, profileMedium?: string | null, profileLarger?: string | null, coverThumb?: string | null, coverMedium?: string | null, coverLarger?: string | null } | null> | null, music?: { __typename?: 'Music', id?: string | null, title?: string | null, playUrl?: string | null, coverThumb?: string | null, coverMedium?: string | null, coverLarge?: string | null, authorName?: string | null, original?: boolean | null, duration?: number | null, album?: string | null } | null, author?: { __typename?: 'Author', id?: string | null, uniqueId?: string | null, nickname?: string | null, avatarThumb?: string | null, avatarMedium?: string | null, avatarLarger?: string | null, signature?: string | null, verified?: boolean | null } | null, video?: { __typename?: 'Video', id?: string | null, height?: number | null, width?: number | null, duration?: number | null, ratio?: string | null, cover?: string | null, originCover?: string | null, dynamicCover?: string | null, playAddr?: string | null, playAddrPublic?: string | null, reflowCover?: string | null, bitrate?: number | null, encodedType?: string | null, format?: string | null, videoQuality?: string | null, encodeUserTag?: string | null, shareCover?: Array<string | null> | null } | null } | null };

export type SearchTiktokQueryVariables = Exact<{
  input: Scalars['String']['input'];
}>;


export type SearchTiktokQuery = { __typename?: 'Query', searchTiktok?: Array<{ __typename?: 'VideoItem', type?: number | null, common?: { __typename?: 'Common', doc_id_str?: string | null } | null, item?: { __typename?: 'Item', id?: string | null, desc?: string | null, originalItem?: boolean | null, officalItem?: boolean | null, secret?: boolean | null, forFriend?: boolean | null, digged?: boolean | null, itemCommentStatus?: number | null, showNotPass?: boolean | null, itemMute?: boolean | null, privateItem?: boolean | null, duetEnabled?: boolean | null, stitchEnabled?: boolean | null, shareEnabled?: boolean | null, authorStats?: { __typename?: 'AuthorStats', followingCount?: number | null, followerCount?: number | null, heartCount?: number | null, videoCount?: number | null, diggCount?: number | null, heart?: number | null } | null, textExtra?: Array<{ __typename?: 'TextExtra', awemeId?: string | null, start?: number | null, end?: number | null, hashtagName?: string | null, hashtagId?: string | null, type?: number | null } | null> | null, stats?: { __typename?: 'Stats', diggCount?: number | null, shareCount?: number | null, commentCount?: number | null, playCount?: number | null, collectCount?: number | null } | null, challenges?: Array<{ __typename?: 'Challenge', id?: string | null, title?: string | null, desc?: string | null, profileThumb?: string | null, profileMedium?: string | null, profileLarger?: string | null, coverThumb?: string | null, coverMedium?: string | null, coverLarger?: string | null } | null> | null, music?: { __typename?: 'Music', id?: string | null, title?: string | null, playUrl?: string | null, coverThumb?: string | null, coverMedium?: string | null, coverLarge?: string | null, authorName?: string | null, original?: boolean | null, duration?: number | null, album?: string | null } | null, author?: { __typename?: 'Author', id?: string | null, uniqueId?: string | null, nickname?: string | null, avatarThumb?: string | null, avatarMedium?: string | null, avatarLarger?: string | null, signature?: string | null, verified?: boolean | null } | null, video?: { __typename?: 'Video', id?: string | null, height?: number | null, width?: number | null, duration?: number | null, ratio?: string | null, cover?: string | null, originCover?: string | null, dynamicCover?: string | null, playAddr?: string | null, playAddrPublic?: string | null, reflowCover?: string | null, bitrate?: number | null, encodedType?: string | null, format?: string | null, videoQuality?: string | null, encodeUserTag?: string | null, shareCover?: Array<string | null> | null } | null } | null }> | null };

export type GetTiktokQueryVariables = Exact<{
  input: Scalars['String']['input'];
}>;


export type GetTiktokQuery = { __typename?: 'Query', getTiktok?: { __typename?: 'VideoItem', type?: number | null, common?: { __typename?: 'Common', doc_id_str?: string | null } | null, item?: { __typename?: 'Item', id?: string | null, desc?: string | null, originalItem?: boolean | null, officalItem?: boolean | null, secret?: boolean | null, forFriend?: boolean | null, digged?: boolean | null, itemCommentStatus?: number | null, showNotPass?: boolean | null, itemMute?: boolean | null, privateItem?: boolean | null, duetEnabled?: boolean | null, stitchEnabled?: boolean | null, shareEnabled?: boolean | null, authorStats?: { __typename?: 'AuthorStats', followingCount?: number | null, followerCount?: number | null, heartCount?: number | null, videoCount?: number | null, diggCount?: number | null, heart?: number | null } | null, textExtra?: Array<{ __typename?: 'TextExtra', awemeId?: string | null, start?: number | null, end?: number | null, hashtagName?: string | null, hashtagId?: string | null, type?: number | null } | null> | null, stats?: { __typename?: 'Stats', diggCount?: number | null, shareCount?: number | null, commentCount?: number | null, playCount?: number | null, collectCount?: number | null } | null, challenges?: Array<{ __typename?: 'Challenge', id?: string | null, title?: string | null, desc?: string | null, profileThumb?: string | null, profileMedium?: string | null, profileLarger?: string | null, coverThumb?: string | null, coverMedium?: string | null, coverLarger?: string | null } | null> | null, music?: { __typename?: 'Music', id?: string | null, title?: string | null, playUrl?: string | null, coverThumb?: string | null, coverMedium?: string | null, coverLarge?: string | null, authorName?: string | null, original?: boolean | null, duration?: number | null, album?: string | null } | null, author?: { __typename?: 'Author', id?: string | null, uniqueId?: string | null, nickname?: string | null, avatarThumb?: string | null, avatarMedium?: string | null, avatarLarger?: string | null, signature?: string | null, verified?: boolean | null } | null, video?: { __typename?: 'Video', id?: string | null, height?: number | null, width?: number | null, duration?: number | null, ratio?: string | null, cover?: string | null, originCover?: string | null, dynamicCover?: string | null, playAddr?: string | null, playAddrPublic?: string | null, reflowCover?: string | null, bitrate?: number | null, encodedType?: string | null, format?: string | null, videoQuality?: string | null, encodeUserTag?: string | null, shareCover?: Array<string | null> | null } | null } | null } | null };

export type SubSearchTiktokSubscriptionVariables = Exact<{
  input: Scalars['String']['input'];
}>;


export type SubSearchTiktokSubscription = { __typename?: 'Subscription', subSearchTiktok?: Array<{ __typename?: 'VideoItem', type?: number | null, common?: { __typename?: 'Common', doc_id_str?: string | null } | null, item?: { __typename?: 'Item', id?: string | null, desc?: string | null, originalItem?: boolean | null, officalItem?: boolean | null, secret?: boolean | null, forFriend?: boolean | null, digged?: boolean | null, itemCommentStatus?: number | null, showNotPass?: boolean | null, itemMute?: boolean | null, privateItem?: boolean | null, duetEnabled?: boolean | null, stitchEnabled?: boolean | null, shareEnabled?: boolean | null, authorStats?: { __typename?: 'AuthorStats', followingCount?: number | null, followerCount?: number | null, heartCount?: number | null, videoCount?: number | null, diggCount?: number | null, heart?: number | null } | null, textExtra?: Array<{ __typename?: 'TextExtra', awemeId?: string | null, start?: number | null, end?: number | null, hashtagName?: string | null, hashtagId?: string | null, type?: number | null } | null> | null, stats?: { __typename?: 'Stats', diggCount?: number | null, shareCount?: number | null, commentCount?: number | null, playCount?: number | null, collectCount?: number | null } | null, challenges?: Array<{ __typename?: 'Challenge', id?: string | null, title?: string | null, desc?: string | null, profileThumb?: string | null, profileMedium?: string | null, profileLarger?: string | null, coverThumb?: string | null, coverMedium?: string | null, coverLarger?: string | null } | null> | null, music?: { __typename?: 'Music', id?: string | null, title?: string | null, playUrl?: string | null, coverThumb?: string | null, coverMedium?: string | null, coverLarge?: string | null, authorName?: string | null, original?: boolean | null, duration?: number | null, album?: string | null } | null, author?: { __typename?: 'Author', id?: string | null, uniqueId?: string | null, nickname?: string | null, avatarThumb?: string | null, avatarMedium?: string | null, avatarLarger?: string | null, signature?: string | null, verified?: boolean | null } | null, video?: { __typename?: 'Video', id?: string | null, height?: number | null, width?: number | null, duration?: number | null, ratio?: string | null, cover?: string | null, originCover?: string | null, dynamicCover?: string | null, playAddr?: string | null, playAddrPublic?: string | null, reflowCover?: string | null, bitrate?: number | null, encodedType?: string | null, format?: string | null, videoQuality?: string | null, encodeUserTag?: string | null, shareCover?: Array<string | null> | null } | null } | null }> | null };


type Properties<T> = Required<{
  [K in keyof T]: z.ZodType<T[K], any, T[K]>;
}>;

type definedNonNullAny = {};

export const isDefinedNonNullAny = (v: any): v is definedNonNullAny => v !== undefined && v !== null;

export const definedNonNullAnySchema = z.any().refine((v) => isDefinedNonNullAny(v));

export function NewTodoDtoSchema(): z.ZodObject<Properties<NewTodoDto>> {
  return z.object({
    text: z.string(),
    userId: z.string()
  })
}

export function UpdateTodoDtoSchema(): z.ZodObject<Properties<UpdateTodoDto>> {
  return z.object({
    done: z.boolean(),
    id: z.string(),
    text: z.string()
  })
}
