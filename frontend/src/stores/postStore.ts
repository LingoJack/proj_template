import { create } from 'zustand'
import { postApi } from '@/api/posts'
import type { Post, CreatePostRequest, UpdatePostRequest } from '@/types/post'

interface PostState {
  posts: Post[]
  total: number
  page: number
  pageSize: number
  loading: boolean
  error: string | null

  fetchPosts: (page?: number, pageSize?: number) => Promise<void>
  createPost: (data: CreatePostRequest) => Promise<void>
  updatePost: (id: number, data: UpdatePostRequest) => Promise<void>
  deletePost: (id: number) => Promise<void>
}

export const usePostStore = create<PostState>((set, get) => ({
  posts: [],
  total: 0,
  page: 1,
  pageSize: 10,
  loading: false,
  error: null,

  fetchPosts: async (page?: number, pageSize?: number) => {
    set({ loading: true, error: null })
    try {
      const p = page ?? get().page
      const ps = pageSize ?? get().pageSize
      const res = await postApi.list(p, ps)
      const { items, total } = res.data.data
      set({ posts: items, total, page: p, pageSize: ps })
    } catch (e) {
      set({ error: (e as Error).message })
    } finally {
      set({ loading: false })
    }
  },

  createPost: async (data) => {
    set({ loading: true, error: null })
    try {
      await postApi.create(data)
      await get().fetchPosts()
    } catch (e) {
      set({ error: (e as Error).message })
    } finally {
      set({ loading: false })
    }
  },

  updatePost: async (id, data) => {
    set({ loading: true, error: null })
    try {
      await postApi.update(id, data)
      await get().fetchPosts()
    } catch (e) {
      set({ error: (e as Error).message })
    } finally {
      set({ loading: false })
    }
  },

  deletePost: async (id) => {
    set({ loading: true, error: null })
    try {
      await postApi.delete(id)
      await get().fetchPosts()
    } catch (e) {
      set({ error: (e as Error).message })
    } finally {
      set({ loading: false })
    }
  },
}))
