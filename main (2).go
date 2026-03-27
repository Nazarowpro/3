package tree

import (
	"testing"
)

// TestBinaryTreeInsert проверяет вставку элементов в дерево
func TestBinaryTreeInsert(t *testing.T) {
	tree := NewBinaryTree[int]()
	
	// Проверяем пустое дерево
	if tree.Root != nil {
		t.Error("Новое дерево должно быть пустым")
	}
	
	// Вставляем корневой элемент
	tree.Insert(50)
	if tree.Root == nil {
		t.Fatal("Корень не должен быть nil после вставки")
	}
	if tree.Root.Value != 50 {
		t.Errorf("Ожидалось значение корня 50, получено %d", tree.Root.Value)
	}
	
	// Вставляем левый потомок
	tree.Insert(30)
	if tree.Root.Left == nil {
		t.Fatal("Левый потомок не должен быть nil")
	}
	if tree.Root.Left.Value != 30 {
		t.Errorf("Ожидалось значение левого потомка 30, получено %d", tree.Root.Left.Value)
	}
	
	// Вставляем правый потомок
	tree.Insert(70)
	if tree.Root.Right == nil {
		t.Fatal("Правый потомок не должен быть nil")
	}
	if tree.Root.Right.Value != 70 {
		t.Errorf("Ожидалось значение правого потомка 70, получено %d", tree.Root.Right.Value)
	}
	
	// Вставляем еще элементы
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(60)
	tree.Insert(80)
	
	// Проверяем структуру дерева
	if tree.Root.Left.Left.Value != 20 {
		t.Errorf("Ожидалось 20, получено %d", tree.Root.Left.Left.Value)
	}
	if tree.Root.Left.Right.Value != 40 {
		t.Errorf("Ожидалось 40, получено %d", tree.Root.Left.Right.Value)
	}
	if tree.Root.Right.Left.Value != 60 {
		t.Errorf("Ожидалось 60, получено %d", tree.Root.Right.Left.Value)
	}
	if tree.Root.Right.Right.Value != 80 {
		t.Errorf("Ожидалось 80, получено %d", tree.Root.Right.Right.Value)
	}
}

// TestBinaryTreeSearch проверяет поиск элементов в дереве
func TestBinaryTreeSearch(t *testing.T) {
	tree := NewBinaryTree[int]()
	
	// Поиск в пустом дереве
	if tree.Search(10) {
		t.Error("Поиск в пустом дереве должен возвращать false")
	}
	
	// Добавляем элементы
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, v := range values {
		tree.Insert(v)
	}
	
	// Проверяем существующие элементы
	for _, v := range values {
		if !tree.Search(v) {
			t.Errorf("Элемент %d должен быть найден", v)
		}
	}
	
	// Проверяем несуществующие элементы
	nonExistent := []int{10, 25, 45, 55, 75, 90}
	for _, v := range nonExistent {
		if tree.Search(v) {
			t.Errorf("Элемент %d не должен быть найден", v)
		}
	}
}

// TestBinaryTreeHeight проверяет вычисление высоты дерева
func TestBinaryTreeHeight(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected int
	}{
		{
			name:     "пустое дерево",
			values:   []int{},
			expected: 0,
		},
		{
			name:     "один элемент",
			values:   []int{50},
			expected: 1,
		},
		{
			name:     "полное дерево",
			values:   []int{50, 30, 70, 20, 40, 60, 80},
			expected: 3,
		},
		{
			name:     "левостороннее дерево",
			values:   []int{50, 40, 30, 20, 10},
			expected: 5,
		},
		{
			name:     "правостороннее дерево",
			values:   []int{10, 20, 30, 40, 50},
			expected: 5,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewBinaryTree[int]()
			for _, v := range tt.values {
				tree.Insert(v)
			}
			
			height := tree.GetHeight()
			if height != tt.expected {
				t.Errorf("Ожидалась высота %d, получена %d", tt.expected, height)
			}
		})
	}
}

// TestBinaryTreeWithStrings проверяет работу дерева со строками
func TestBinaryTreeWithStrings(t *testing.T) {
	tree := NewBinaryTree[string]()
	
	strings := []string{"banana", "apple", "cherry", "date", "fig"}
	for _, s := range strings {
		tree.Insert(s)
	}
	
	// Проверяем поиск
	if !tree.Search("apple") {
		t.Error("Элемент 'apple' должен быть найден")
	}
	if !tree.Search("cherry") {
		t.Error("Элемент 'cherry' должен быть найден")
	}
	if tree.Search("grape") {
		t.Error("Элемент 'grape' не должен быть найден")
	}
	
	// Проверяем высоту
	height := tree.GetHeight()
	if height <= 0 {
		t.Errorf("Высота дерева должна быть положительной, получена %d", height)
	}
}

// TestBinaryTreeTraversals проверяет обходы дерева
func TestBinaryTreeTraversals(t *testing.T) {
	tree := NewBinaryTree[int]()
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, v := range values {
		tree.Insert(v)
	}
	
	// Просто проверяем, что обходы не вызывают паники
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Обход дерева вызвал панику: %v", r)
		}
	}()
	
	tree.InOrder()
	tree.PreOrder()
	tree.PostOrder()
}

// TestDuplicateValues проверяет вставку дубликатов
func TestDuplicateValues(t *testing.T) {
	tree := NewBinaryTree[int]()
	
	tree.Insert(50)
	tree.Insert(50) // Вставляем дубликат
	
	// Дубликат должен быть вставлен в правое поддерево
	if tree.Root.Right == nil {
		t.Error("Дубликат должен быть вставлен в правое поддерево")
	}
	if tree.Root.Right.Value != 50 {
		t.Errorf("Ожидалось значение 50, получено %d", tree.Root.Right.Value)
	}
}

// TestBinaryTreeLarge проверяет работу с большим количеством элементов
func TestBinaryTreeLarge(t *testing.T) {
	tree := NewBinaryTree[int]()
	
	// Добавляем 100 элементов
	for i := 0; i < 100; i++ {
		tree.Insert(i)
	}
	
	// Проверяем поиск
	for i := 0; i < 100; i++ {
		if !tree.Search(i) {
			t.Errorf("Элемент %d не найден", i)
		}
	}
	
	// Высота должна быть 100 (так как вставляем последовательные числа)
	height := tree.GetHeight()
	if height != 100 {
		t.Errorf("Ожидалась высота 100, получена %d", height)
	}
}