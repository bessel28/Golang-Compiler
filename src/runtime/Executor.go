package runtime

//
//func (s *BlockStatement) Execute(env *Environment) error {
//	local := env.NewChild()
//	for _, stmt := range s.Statements {
//		if err := stmt.Execute(local); err != nil {
//			return err
//		}
//	}
//	return nil
//}

//func (e ExpressionStatement) Execute(env *Environment) error {
//_, err := e.Expression.Eval(env)
//return err
//}

//func (s *VariableAssignmentStatement) Execute(env *runtime.Environment) error {
//val, err := s.Value.Eval(env)
//if err != nil {
//return err
//}
//return env.Assign(s.Name, &val)
//}
//
//func (s *VariableDeclarationStatement) Execute(env *runtime.Environment) error {
//val, err := s.Value.Eval(env)
//if err != nil {
//return err
//}
//return env.Define(s.Name, &val)
//}
