Your goal is to implement a type system for a custom Toy Language. This system must handle primitive types, generics, nested tuples, and function signatures. You will implement the core data structures and a Type Inference engine that substitutes generics with concrete types.

Type Definitions:

Primitives: Lowercase strings like int, float, str, bool
Generics: Uppercase letters followed by numbers, e.g., T1, T2
Tuples: Comma-separated types inside brackets, which can be nested. Example: [int, [T1, str]].
Functions: Defined by a list of parameter types and a single return type. Syntax: [param1, param2] -> returnType.
Part 1: Implement to_str in Node and Function.

Node Class: Represents a type node. It can be a leaf (primitive/generic) or a tuple (list of child nodes).

class Node:
    def  __init__(self, node_type):

        If node_type is a string: It is a primitive or generic.

        If node_type is a list: It is a tuple containing other Node objects.

    def to_str(self):

        Primitives/Generics: Return the string name (e.g., "int").

        Tuples: Return bracketed, comma-separated types (e.g., "[int,T1]").

class Function:
    def    __init__(self, parameters, output_type):
        Represents a function signature.
        parameters: A List[Node] objects.

        output_type: A single Node object.

    def to_str(self):

        Format: (param1,param2,...) -> returnType.

        Example: (int,T1) -> [T1,str]
Part 2: Implement a function get_return_type(parameters, function) that determines the concrete return type of a function based on provided arguments.

def get_return_type(parameters: List[Node], function: Function) -> Node:
    pass
Requirements:

Generic Resolution: Build a mapping (substitution table) by comparing the Function's expected parameters to the actual parameters provided.
Substitution: Recursively replace all generics in the function's output_type with the concrete types found during resolution.
Error Handling:
Argument Count Mismatch: Raise an error if the number of arguments doesn't match.
Type Mismatch: Raise an error if a concrete type (e.g., int) is expected but a different type (e.g., str) is provided.
Generic Conflict: Raise an error if the same generic (e.g., T1) is bound to two different concrete types in the same call.
Test Examples:

Example 1: Basic Substitution

    Function: [T1, T2, int, T1] -> [T1, T2]
    Arguments: [int, str, int, int]
    Logic: T1 maps to int, T2 maps to str.
    Result: [int, str]

Example 2: Nested Tuples & Complex Generics

    Function: [[T1, float], T2, T3] -> [T3, T1]
    Arguments: [[str, float], [int, str], int]
    Logic: * T1 is extracted from the first tuple as str.
        T2 maps to the tuple [int, str].
        T3 maps to int.

    Result: [int, str]

Example 3: Conflict Error
    Function: [T1, T1] -> T1
    Arguments: [int, str]
    Error: T1 cannot be both int and str.