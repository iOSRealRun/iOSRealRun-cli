"""
parseRoute.py
parse the route from the input
"""
def split(inp):
    import ast
    tmp = ast.literal_eval("[{}]".format(inp))
    for i in tmp:
        i["lat"] = float(i["lat"])
        i["lng"] = float(i["lng"])
    return tmp


