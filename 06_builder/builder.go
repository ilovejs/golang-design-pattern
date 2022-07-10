package builder

//Builder 是生成器接口
type Builder interface {
	Part1()
	Part2()
	Part3()
}

// Director
type Director struct {
	builder Builder // todo: 嵌入一个接口
}

// 返回结构体的指针，而不是接口
func NewDirector(
	builder Builder, // take a builder interface, so we can replace it with builder 1 and 2.
) *Director {
	return &Director{
		builder: builder,
	}
}

// director 掌握部件的需求
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

// Builder1
type Builder1 struct {
	result string
}

// builder 知道部件的制作过程
func (b *Builder1) Part1() {
	b.result += "1"
}

func (b *Builder1) Part2() {
	b.result += "2"
}

func (b *Builder1) Part3() {
	b.result += "3"
}

func (b *Builder1) GetResult() string {
	return b.result
}

// Builder2
type Builder2 struct {
	result int // diff type
}

func (b *Builder2) Part1() {
	b.result += 1
}

func (b *Builder2) Part2() {
	b.result += 2
}

func (b *Builder2) Part3() {
	b.result += 3
}

func (b *Builder2) GetResult() int {
	return b.result
}
