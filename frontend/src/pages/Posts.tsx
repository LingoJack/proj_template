import { useEffect, useState } from 'react'
import { usePostStore } from '@/stores/postStore'
import type { CreatePostRequest } from '@/types/post'

export default function Posts() {
  const { posts, total, page, loading, error, fetchPosts, createPost, deletePost } = usePostStore()
  const [showForm, setShowForm] = useState(false)
  const [form, setForm] = useState<CreatePostRequest>({ title: '', content: '' })

  useEffect(() => {
    fetchPosts()
  }, [fetchPosts])

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    await createPost(form)
    setForm({ title: '', content: '' })
    setShowForm(false)
  }

  return (
    <div className="max-w-3xl mx-auto">
      <div className="flex items-center justify-between mb-4">
        <h2 className="text-2xl font-bold text-gray-800">文章</h2>
        <button
          onClick={() => setShowForm(!showForm)}
          className="px-4 py-2 bg-blue-600 text-white text-sm rounded-md hover:bg-blue-700 transition-colors"
        >
          {showForm ? '取消' : '新建'}
        </button>
      </div>

      {error && (
        <div className="mb-4 p-3 bg-red-50 border border-red-200 rounded-md text-red-700 text-sm">
          {error}
        </div>
      )}

      {showForm && (
        <form onSubmit={handleSubmit} className="mb-6 bg-white p-4 rounded-lg border border-gray-200 space-y-3">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">标题</label>
            <input
              type="text"
              value={form.title}
              onChange={(e) => setForm({ ...form, title: e.target.value })}
              required
              className="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">内容</label>
            <textarea
              value={form.content}
              onChange={(e) => setForm({ ...form, content: e.target.value })}
              required
              rows={3}
              className="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
          <button
            type="submit"
            disabled={loading}
            className="px-4 py-2 bg-blue-600 text-white text-sm rounded-md hover:bg-blue-700 disabled:opacity-50 transition-colors"
          >
            {loading ? '提交中...' : '提交'}
          </button>
        </form>
      )}

      {loading && !posts.length ? (
        <p className="text-gray-500 text-sm">加载中...</p>
      ) : posts.length === 0 ? (
        <p className="text-gray-500 text-sm">暂无文章</p>
      ) : (
        <div className="space-y-3">
          {posts.map((post) => (
            <div key={post.id} className="bg-white p-4 rounded-lg border border-gray-200">
              <div className="flex items-start justify-between">
                <div>
                  <h3 className="font-medium text-gray-800">{post.title}</h3>
                  <p className="text-sm text-gray-500 mt-1">{post.content}</p>
                </div>
                <button
                  onClick={() => deletePost(post.id)}
                  className="text-xs text-red-500 hover:text-red-700 transition-colors"
                >
                  删除
                </button>
              </div>
            </div>
          ))}
        </div>
      )}

      {total > 0 && (
        <div className="mt-4 flex items-center justify-between text-sm text-gray-500">
          <span>共 {total} 条</span>
          <div className="flex gap-2">
            <button
              onClick={() => fetchPosts(page - 1)}
              disabled={page <= 1}
              className="px-3 py-1 border border-gray-300 rounded-md disabled:opacity-50 hover:bg-gray-100 transition-colors"
            >
              上一页
            </button>
            <span className="px-3 py-1">第 {page} 页</span>
            <button
              onClick={() => fetchPosts(page + 1)}
              disabled={page * 10 >= total}
              className="px-3 py-1 border border-gray-300 rounded-md disabled:opacity-50 hover:bg-gray-100 transition-colors"
            >
              下一页
            </button>
          </div>
        </div>
      )}
    </div>
  )
}
