import React, { useEffect, useState } from "react";
import {
  PlusCircle,
  CheckCircle2,
  Circle,
  Trash2,
  Search,
  ListTodo,
} from "lucide-react";

interface Todo {
  ID: number;
  title: string;
  done: boolean;
  created_at?: string | null;
  updated_at?: string | null;
  deleted_at?: string | null;
}

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [newTodo, setNewTodo] = useState("");
  const [searchQuery, setSearchQuery] = useState("");

  useEffect(() => {
    // const controller = new AbortController();

    async function fetchTodos() {
      try {
        const response = await fetch("http://localhost:8080/todos/");
        if (!response.ok) {
          throw new Error(`HTTP error! Status: ${response.status}`);
        }
        const data = await response.json();
        console.log(data);
        setTodos(data);
      } catch (error) {
        console.log("Fetch error:", error);
      }
    }

    fetchTodos();
  }, []);

  const addTodo = (e: React.FormEvent) => {
    e.preventDefault();
    if (newTodo.trim()) {
      // setTodos([
      //   ...todos,
      //   { ID: crypto.randomUUID(), title: newTodo.trim(), done: false },
      // ]);
      // setNewTodo("");
    }
  };

  const toggleTodo = (id: number) => {
    setTodos(
      todos.map((todo) =>
        todo.ID === id ? { ...todo, done: !todo.done } : todo
      )
    );
  };

  const deleteTodo = (id: number) => {
    setTodos(todos.filter((todo) => todo.ID !== id));
  };

  const filteredTodos = todos?.filter((todo) =>
    todo?.title.toLowerCase().includes(searchQuery?.toLowerCase())
  );

  const completedTodos = todos.filter((todo) => todo.done).length;

  return (
    <div className="min-h-screen bg-gradient-to-br from-purple-50 to-blue-50 py-8 px-4">
      <div className="max-w-2xl mx-auto">
        {/* Header */}
        <div className="flex items-center gap-3 mb-8">
          <ListTodo className="w-8 h-8 text-purple-600" />
          <h1 className="text-3xl font-bold text-gray-800">My Tasks</h1>
        </div>

        {/* Add Todo Form */}
        <form onSubmit={addTodo} className="mb-8">
          <div className="flex gap-2">
            <input
              type="text"
              value={newTodo}
              onChange={(e) => setNewTodo(e.target.value)}
              placeholder="Add a new task..."
              className="flex-1 px-4 py-3 rounded-lg border border-gray-200 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent transition"
            />
            <button
              type="submit"
              className="px-6 py-3 bg-purple-600 text-white rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2 transition flex items-center gap-2"
            >
              <PlusCircle className="w-5 h-5" />
              Add Task
            </button>
          </div>
        </form>

        {/* Search and Stats */}
        <div className="mb-6 flex items-center justify-between">
          <div className="relative flex-1 max-w-xs">
            <Search className="w-5 h-5 absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" />
            <input
              type="text"
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              placeholder="Search tasks..."
              className="pl-10 pr-4 py-2 w-full rounded-lg border border-gray-200 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent transition"
            />
          </div>
          <div className="text-sm text-gray-600">
            Completed: {completedTodos} of {todos.length}
          </div>
        </div>

        {/* Todo List */}
        <div className="bg-white rounded-xl shadow-sm overflow-hidden">
          {filteredTodos.length === 0 ? (
            <div className="p-8 text-center text-gray-500">
              {todos.length === 0
                ? "No tasks yet. Add your first task!"
                : "No matching tasks found."}
            </div>
          ) : (
            <ul className="divide-y divide-gray-100">
              {filteredTodos.map((todo) => (
                <li
                  key={todo.ID}
                  className="flex items-center gap-4 px-6 py-4 hover:bg-gray-50 transition group"
                >
                  <button
                    onClick={() => toggleTodo(todo.ID)}
                    className="flex-shrink-0 focus:outline-none"
                  >
                    {todo.done ? (
                      <CheckCircle2 className="w-6 h-6 text-green-500" />
                    ) : (
                      <Circle className="w-6 h-6 text-gray-400" />
                    )}
                  </button>
                  <span
                    className={`flex-1 ${
                      todo.done ? "text-gray-400 line-through" : "text-gray-700"
                    }`}
                  >
                    {todo.title}
                  </span>
                  <button
                    onClick={() => deleteTodo(todo.ID)}
                    className="opacity-0 group-hover:opacity-100 focus:opacity-100 p-2 text-gray-400 hover:text-red-500 transition focus:outline-none"
                  >
                    <Trash2 className="w-5 h-5" />
                  </button>
                </li>
              ))}
            </ul>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;
