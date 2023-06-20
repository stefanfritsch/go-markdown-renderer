package parse_dir

type IEntry interface {
	Priority(...int) int          // Used for sorting. Higher priority == listed first
	DisplayName(...string) string // How should the name be presented to the user?
	BaseName() string             // The filename or dirname from the last slash
	DirName() string              // The dirname up to the last slash
	Path(...string) string        // The full path to the element
	Parent(...IDir) IDir          // Parent of the current entry
	Level(...int) int             // How deep are we in the tree? root: 0, files and subdirs in root: 1, etc.
	Content(...string) string     // Contents of the object
}

type IFile interface {
	IEntry
}

type IDir interface {
	IEntry
	Subdirs(set ...map[string]IDir) map[string]IDir // subdirectories of this dir
	SubdirOrder(set ...[]string) []string           // The keys of Subdirs() in the right order for use
	Files(set ...map[string]IFile) map[string]IFile // files in this dir
	FileOrder(set ...[]string) []string             // The keys of Files() in the right order for use
	Children() []IEntry                             // subdirectories + files
}

type Entry struct {
	priority    int    // Used for sorting. Higher priority == listed first
	displayName string // How should the name be presented to the user?
	path        string // The full path to the element
	parent      IDir   // the parent directory of this entry
	level       int    // How deep are we in the tree? root: 0, files and subdirs in root: 1, etc.
	content     string // Contents of the object
}

func (e *Entry) Priority(set ...int) int {
	if len(set) > 0 {
		e.priority = set[0]
	}
	return e.priority
}
func (e *Entry) DisplayName(set ...string) string {
	if len(set) > 0 {
		e.displayName = set[0]
	}
	return e.displayName
}
func (e *Entry) BaseName() string {
	return GetBaseName(e.path)
}
func (e *Entry) DirName() string {
	return GetDirName(e.path)
}

func (e *Entry) Path(set ...string) string {
	if len(set) > 0 {
		e.path = set[0]
	}
	return e.path
}

func (e *Entry) Parent(set ...IDir) IDir {
	if len(set) > 0 {
		e.parent = set[0]
	}

	return e.parent
}

func (e *Entry) Level(set ...int) int {
	if len(set) > 0 {
		e.level = set[0]
	}

	return e.level
}

func (e *Entry) Content(set ...string) string {
	if len(set) > 0 {
		e.content = set[0]
	}

	return e.content
}

type File struct {
	Entry
}

type Dir struct {
	Entry

	subdirs     map[string]IDir
	subdirOrder []string
	files       map[string]IFile
	fileOrder   []string
}

func (d *Dir) Subdirs(set ...map[string]IDir) map[string]IDir {
	if len(set) > 0 {
		setdirs := set[0]

		for k, v := range setdirs {
			d.subdirs[k] = v
		}
	}
	return d.subdirs
}

func (d *Dir) SubdirOrder(set ...[]string) []string {
	if len(set) > 0 {
		d.subdirOrder = set[0]
	}

	return d.subdirOrder
}

func (d *Dir) Files(set ...map[string]IFile) map[string]IFile {
	if len(set) > 0 {
		setfiles := set[0]

		for k, v := range setfiles {
			d.files[k] = v
		}
	}
	return d.files
}

func (d *Dir) FileOrder(set ...[]string) []string {
	if len(set) > 0 {
		d.fileOrder = set[0]
	}

	return d.fileOrder
}

func (d *Dir) Children() []IEntry {
	var children []IEntry

	for _, subdir := range d.subdirs {
		children = append(children, subdir)
	}

	return children
}
