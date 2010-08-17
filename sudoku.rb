#!/usr/bin/env ruby

class Cell

    attr_reader :board, :row, :col, :possibles, :ref
    
    def initialize(board, row, col, ref)
        @board = board
        @row = row
        @col = col
        @ref = ref
        @possibles = '123456789'
    end
    
    def box()
        (@row / 3) * 3 + (@col / 3)
    end
    
    def has?(possible)
        @possibles.length > 1 and @possibles.include? possible
    end
    
    def solve(value)
        @possibles = value
    end
    
    def remove(value)
        found = false
        if @possibles.length > 1
            found = @possibles.sub!(value, '')
#            puts "removed #{value} from #{self}" if found
            if @possibles.length == 1
#                puts "#{self} solved"
                @board.solve(self, @possibles)
            end
        end
        found
    end
    
    def remove_possibles(values)
        found = false        
        values.scan(/./).each { |value| found = true if remove(value) }
        found
    end
    
    def to_s()
        "Cell #{@ref} #{@row}:#{@col}=#{@possibles}"
    end
    
end 

class Board

    def initialize(data)
        @cells = []
        i = 0
        (0..8).each { |row| (0..8).each { |col| @cells << Cell.new(self, row, col, i); i += 1 }}
        i = 0
        (0..8).each { |row| (0..8).each { |col| 
            value = data[i,1]
            solve(cell(row, col), value) if '123456789'.include? value 
            i += 1
        }}
            
    end
    
    def to_s()
        rows = []
        (0..8).each do |row|
            rows << cells_by_row(row).each.map { |cell| cell.possibles }
        end
        (0..8).each do |col|
            max = 0
            (0..8).each do |row|
                max = rows[row][col].length if rows[row][col].length > max
            end
            (0..8).each do |row|
                rows[row][col] = rows[row][col] + " " * (max - rows[row][col].length)
            end
        end
        s = []
        s << '=' * 50
        i = 0
        rows.each do |row|
            i += 1
            s << "#{row.join("|")}"
            s << '-' * 50 if i == 3 or i == 6
        end
        s << '=' * 50
        s.join("\n")
    end
    
    def solve(solved_cell, value)
        solved_cell.solve(value)
        @cells.each { |cell| cell.remove(value) if cell.row == solved_cell.row or cell.col == solved_cell.col or cell.box == solved_cell.box }
    end
    
    def cell(row, col)
        @cells[row * 9 + col]
    end
    
    def cells_by_row(row)
        @cells.select { |cell| cell.row == row }
    end

    def cells_by_col(col)
        @cells.select { |cell| cell.col == col }
    end

    def cells_by_box(box)
        @cells.select { |cell| cell.box == box }
    end
    
    def cell_by_ref(ref)
        @cells.select { |cell| cell.ref == ref }[0]
    end
    
    def singles
        puts "=== Singles ==="
        found = false
        
        (0..8).each do |row|
            "123456789".scan(/./) do |possible|
                cells = cells_by_row(row).select { |cell| cell.has?(possible) }
                if cells.length == 1
                    puts "Single #{possible} in row #{row}"
                    solve(cells[0], possible)
                    found = true
                end
            end
        end
        
        (0..8).each do |col|
            "123456789".scan(/./) do |possible|
                cells = cells_by_col(col).select { |cell| cell.has?(possible) }
                if cells.length == 1
                    puts "Single #{possible} in col #{col}"
                    solve(cells[0], possible)
                    found = true
                end
            end
        end
        
        (0..8).each do |box|
            "123456789".scan(/./) do |possible|
                cells = cells_by_box(box).select { |cell| cell.has?(possible) }
                if cells.length == 1
                    puts "Single #{possible} in box #{box}"
                    solve(cells[0], possible)
                    found = true
                end
            end
        end
        
        found
    end

    def naked_pairs
        # http://www.sudokuwiki.org/Naked_Candidates#NP
        puts "=== Naked Pairs ==="
        found = false
        row = col = 0
        
        (0..8).each do |row|
            pairs = cells_by_row(row).select { |cell| cell.possibles.length == 2}
            while pairs.length >= 2
                first = pairs.pop
                pairs.each do |second|
                    if first.possibles == second.possibles
                        cells_by_row(row).each do |cell|
                            if cell != first and cell != second and cell.remove_possibles(first.possibles)
                                puts "Naked Pair #{first.possibles} on row #{row}"
                                found = true
                            end
                        end
                    end
                end
            end
        end
        
        (0..8).each do |col|
            pairs = cells_by_col(col).select { |cell| cell.possibles.length == 2}
            while pairs.length >= 2
                first = pairs.pop
                pairs.each do |second|
                    if first.possibles == second.possibles
                        cells_by_col(col).each do |cell|
                            if cell != first and cell != second and cell.remove_possibles(first.possibles)
                                puts "Naked Pair #{first.possibles} on col #{col}"
                                found = true
                            end
                        end
                    end
                end
            end
        end
        
        (0..8).each do |box|
            pairs = cells_by_box(box).select { |cell| cell.possibles.length == 2}
            while pairs.length >= 2
                first = pairs.pop
                pairs.each do |second|
                    if first.possibles == second.possibles
                        cells_by_box(box).each do |cell|
                            if cell != first and cell != second and cell.remove_possibles(first.possibles)
                                puts "Naked Pair #{first.possibles} on box #{box}"
                                found = true
                            end
                        end
                    end
                end
            end
        end
        
        found        
    end
        
    def combinations(array, r)
        n = array.length
        indices = (0...r).to_a
        final = (n - r...n).to_a
        while indices != final
            yield indices.map {|k| array[k]}
            i = r - 1
            while indices[i] == n - r + i
                i -= 1
            end
            indices[i] += 1
            (i + 1...r).each do |j|
                indices[j] = indices[i] + j - i
            end
        end
        yield indices.map {|k| array[k]}
    end

        
        
    def naked_cells( cells )
        found = false
        
        cells = cells.select { |cell| cell.possibles.length > 1 }
        possibles = cells.map { |cell| cell.possibles }.join.split(//).uniq.sort
        
        combos = []
        (2...possibles.length).each do |r|
            combinations(possibles, r) { |x| combos << x }
        end
        
        combos.each do |combo|
            matches = cells.select { |cell| (cell.possibles.split(//) - combo).length == 0 }
            if matches.length == combo.length
                remove = combo.join
                cells.each do |cell|
                    if not matches.include? cell
                        found = true if cell.remove_possibles(remove)
                    end
                end
                if found
                    where = matches.map { |cell| "#{cell.row}:#{cell.col}" }.join(" ")
                    puts "Naked #{remove} found in #{where}"
                    return true
                end
            end
        end
        
        found
    end
    
    def nakeds
        # http://www.sudokuwiki.org/Naked_Candidates#NP
        puts "=== Nakeds ==="
        (0..8).each { |row| return true if naked_cells(cells_by_row(row)) }
        (0..8).each { |col| return true if naked_cells(cells_by_col(col)) }
        (0..8).each { |box| return true if naked_cells(cells_by_box(box)) }
        false
    end
    
    def pointing_pairs
        # http://www.sudokuwiki.org/Intersection_Removal#IR
        puts "=== Pointing Pairs ==="
        found = false
        row = col = 0
        
        (0..8).each do |box|
            "123456789".scan(/./).each do |possible|
                pairs = cells_by_box(box).select { |cell| cell.has? possible}
                if pairs.length == 2 or pairs.length == 3
                    if pairs[0].row == pairs[1].row and pairs[0].row == pairs[-1].row
                        cells_by_row(pairs[0].row).each do |cell| 
                            if not pairs.include? cell and cell.remove(possible) 
                                found = true
                            end
                        end
                    end
                    if pairs[0].col == pairs[1].col and pairs[0].col == pairs[-1].col
                        cells_by_col(pairs[0].col).each do |cell| 
                            if not pairs.include? cell and cell.remove(possible) 
                                found = true
                            end
                        end
                    end
                    if found
                        puts "Pointing Pair #{possible} on box #{box}"
                        return true
                    end
                end
            end
        end
        
        found        
    end
        
    def box_line_reduction
        # http://www.sudokuwiki.org/Intersection_Removal#LBR
        puts "=== Box-Line Reduction ==="
        found = false
        row = col = 0
        
        (0..8).each do |row|
            "123456789".scan(/./).each do |possible|
                cells = cells_by_row(row).select { |cell| cell.has? possible }
                boxes = cells.map { |cell| cell.box}.uniq
                if boxes.length == 1
                    cells_by_box(boxes[0]).each do |cell| 
                        if not cells.include? cell and cell.remove(possible) 
                            found = true
                        end
                    end
                end
                if found
                    puts "Box Line reduction of #{possible} on row #{row}"
                    return true
                end
            end
        end
        
        (0..8).each do |col|
            "123456789".scan(/./).each do |possible|
                cells = cells_by_col(col).select { |cell| cell.has? possible }
                boxes = cells.map { |cell| cell.box}.uniq
                if boxes.length == 1
                    cells_by_box(boxes[0]).each do |cell| 
                        if not cells.include? cell and cell.remove(possible) 
                            found = true
                        end
                    end
                end
                if found
                    puts "Box Line reduction of #{possible} on col #{col}"
                    return true
                end
            end
        end
        
        found        
    end
        
    def x_wing
        # http://www.sudokuwiki.org/X_Wing_Strategy
        puts "=== X-Wing ==="
        found = false
        row = col = 0

        "123456789".scan(/./).each do |possible|
            rows = []        
            (0..8).each do |row|
                cols = [possible, row] << cells_by_row(row).select { |cell| cell.has? possible }.map { |cell| cell.col }
                cols.flatten!
                rows << cols if cols.length == 4
            end
            
            while rows.length > 1
                first = rows.pop
                rows.each do |second|
                    if first[2] == second[2] and first[3] == second[3]
                        x_rows = [first[1], second[1]]
                        [first[2], first[3]].each do |col|
                            cells_by_col(col).each do |cell|
                                if not x_rows.include? cell.row and cell.remove(possible) 
                                    found = true
                                end
                            end
                        end
                        if found
                            puts "X-Wing for #{possible} in rows #{x_rows}"
                            return true
                        end
                    end
                end
            end
            
        end
        
        "123456789".scan(/./).each do |possible|
            cols = []        
            (0..8).each do |col|
                rows = [possible, col] << cells_by_col(col).select { |cell| cell.has? possible }.map { |cell| cell.row }
                rows.flatten!
                cols << rows if rows.length == 4
            end
            
            while cols.length > 1
                first = cols.pop
                cols.each do |second|
                    if first[2] == second[2] and first[3] == second[3]
                        x_cols = [first[1], second[1]]
                        [first[2], first[3]].each do |row|
                            cells_by_row(row).each do |cell|
                                if not x_cols.include? cell.col and cell.remove(possible) 
                                    found = true
                                end
                            end
                        end
                        if found
                            puts "X-Wing for #{possible} in cols #{x_cols}"
                            return true
                        end
                    end
                end
            end
            
        end
        
        found        
    end
        
    def swordfish
        # http://www.sudokuwiki.org/Sword_Fish_Strategy
        puts "=== Swordfish ==="
        found = false
        matches = []

        "123456789".scan(/./).each do |possible|
            rows = []        
            (0..8).each do |row|
                cols = [possible, row] + cells_by_row(row).select { |cell| cell.has? possible }.map { |cell| cell.col }
                rows << cols if cols.length == 4 || cols.length == 5
            end
            
            rows.each do |row|
                if row.length == 5
                    triple = [row] + rows.select { |match| match != row and (match[2..-1] - row[2..-1]).length == 0 }
                    if triple.length == 3
                        triple_cols = row[2..-1]
                        triple_rows = triple.map { |x| x[1] }
                        triple_cols.each do |col|
                            cells_by_col(col).each do |cell|
                                if not triple_rows.include? cell.row 
                                    found = true if cell.remove(possible)
                                end
                            end
                        end
                        if found
                            puts "triple=#{triple.inspect}"
                            puts "rows=#{triple_rows.inspect} cols=#{triple_cols.inspect}"
                            return true
                        end
                    end
                end
            end
                
            
        end
        
        "123456789".scan(/./).each do |possible|
            cols = []        
            (0..8).each do |col|
                rows = [possible, col] + cells_by_col(col).select { |cell| cell.has? possible }.map { |cell| cell.row }
                cols << rows if rows.length == 4 || rows.length == 5
            end
            
            cols.each do |col|
                if col.length == 5
                    triple = [col] + cols.select { |match| match != col and (match[2..-1] - col[2..-1]).length == 0 }
                    if triple.length == 3
                        triple_rows = col[2..-1]
                        triple_cols = triple.map { |x| x[1] }
                        triple_rows.each do |row|
                            cells_by_row(row).each do |cell|
                                if not triple_cols.include? cell.col 
                                    found = true if cell.remove(possible)
                                end
                            end
                        end
                        if found
                            puts "triple=#{triple.inspect}"
                            puts "cols=#{triple_cols.inspect} rows=#{triple_rows.inspect}"
                            return true 
                        end
                    end
                end
            end
                
            
        end
        
        found        
    end
    
    def solved
        @cells.all? { |cell| cell.possibles.length == 1 }
    end
        
end

line = 0
File.new("data.txt", "r").each do |data|
    exit if data.length < 81
    line += 1
    puts "=== Populating board #{line} ==="
    puts data
    board = Board.new(data)
    running = true
    while running
        if board.singles
        elsif board.nakeds
        elsif board.pointing_pairs
        elsif board.box_line_reduction
        elsif board.x_wing
        elsif board.swordfish
        else 
            running = false
        end
        puts board.to_s
    end
    exit if not board.solved
end


